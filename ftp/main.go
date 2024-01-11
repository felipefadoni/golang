package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

func main() {
	// Defina flags com os tipos correspondentes
	host := flag.String("host", "", "Especifique o host")
	user := flag.String("user", "", "Especifique o usuário")
	pass := flag.String("pass", "", "Especifique o password")
	path := flag.String("path", "", "Especifique o path")

	// Analise os argumentos da linha de comando
	flag.Parse()

	// Verifique se os parâmetros obrigatórios foram fornecidos
	if *host == "" || *user == "" || *pass == "" || *path == "" {
		fmt.Println("Erro: Os parâmetros --host, --user, --pass e --path são obrigatórios.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Conecte-se ao servidor FTP
	ftpClient, err := ftp.Dial(*host)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor FTP:", err)
		os.Exit(1)
	}
	defer ftpClient.Quit()

	// Faça login no servidor FTP
	err = ftpClient.Login(*user, *pass)
	if err != nil {
		fmt.Println("Erro ao fazer login no servidor FTP:", err)
		os.Exit(1)
	}

	// Mude para o diretório especificado
	err = ftpClient.ChangeDir(*path)
	if err != nil {
		fmt.Println("Erro ao mudar para o diretório especificado:", err)
		os.Exit(1)
	}

	// Liste os arquivos no diretório
	files, err := ftpClient.List("")
	if err != nil {
		fmt.Println("Erro ao listar os arquivos:", err)
		os.Exit(1)
	}

	// Imprima os nomes dos arquivos
	fmt.Println("Arquivos no diretório:")
	for _, file := range files {
		fmt.Println(file.Name)
	}
}
