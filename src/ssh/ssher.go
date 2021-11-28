package ssh

import (
	"encoding/json"
	"fmt"
	"github.com/faradey/madock/src/paths"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"strings"
	"syscall"
	"time"
)

var passwd string

type RemoteDbStruct struct {
	Host           string `json:"host"`
	Dbname         string `json:"dbname"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Active         string `json:"active"`
	Model          string `json:"model"`
	Engine         string `json:"engine"`
	InitStatements string `json:"initStatements"`
}

func RunCommand(conn *ssh.Client, cmd string) string {
	sess, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	out, err := sess.CombinedOutput(cmd)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return string(out)
}

func DbDump(conn *ssh.Client, remoteDir string) {
	defer conn.Close()
	result := RunCommand(conn, "php -r \"\\$r1 = include('"+remoteDir+"/app/etc/env.php'); echo json_encode(\\$r1[\\\"db\\\"][\\\"connection\\\"][\\\"default\\\"]);\"")
	if len(result) > 2 {
		dbAuthData := RemoteDbStruct{}
		err := json.Unmarshal([]byte(result), &dbAuthData)
		if err != nil {
			fmt.Println(err)
		}
		curDateTime := time.Now().Format("2006-01-02_15-04-05")
		dumpName := "dump_db_" + curDateTime + ".sql.gz"
		result = RunCommand(conn, "mysqldump -u "+dbAuthData.Username+" -p"+dbAuthData.Password+" --single-transaction --quick --lock-tables=false --no-tablespaces --triggers "+dbAuthData.Dbname+" | sed -e 's/DEFINER[ ]*=[ ]*[^*]*\\*/\\*/' | gzip > "+remoteDir+"/"+dumpName)
		sc, err := sftp.NewClient(conn)
		if err != nil {
			log.Fatal(err)
		}
		defer sc.Close()
		execPath := paths.GetExecDirPath()
		projectName := paths.GetRunDirName()
		err = downloadFile(sc, remoteDir+"/"+dumpName, execPath+"/projects/"+projectName+"/backup/db/"+dumpName)
		if err != nil {
			log.Fatal(err)
		}
		result = RunCommand(conn, "rm "+remoteDir+"/"+dumpName)
	} else {
		fmt.Println("Failed to get database authentication data")
	}
}

func Connect(authType, keyPath, pswrd, host, port, username string) *ssh.Client {
	config := &ssh.ClientConfig{
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if authType == "password" {
		config.Auth = []ssh.AuthMethod{
			ssh.Password(pswrd),
		}
	} else {
		config.Auth = []ssh.AuthMethod{
			publicKey(keyPath),
		}
	}

	conn, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func Disconnect(conn *ssh.Client) {
	conn.Close()
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		if passwd == "" {
			fmt.Println("Input your password for ssh key:")
			var sentence []byte
			sentence, err = terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatalln(err)
			}
			passwd = strings.TrimSpace(string(sentence))
		}
		signer, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(passwd))
		if err != nil {
			log.Fatal(err)
		}
	}

	return ssh.PublicKeys(signer)
}
