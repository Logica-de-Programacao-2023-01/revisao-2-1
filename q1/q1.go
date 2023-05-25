package q1

//Você está trabalhando em um projeto de gerenciamento de uma escola. O sistema precisa armazenar informações sobre os alunos, incluindo seu nome, idade e as matérias em que estão matriculados, juntamente com suas respectivas notas. Você decidiu usar structs e mapas para representar essas informações.
//
//No entanto, você descobriu que existem dois conjuntos de dados diferentes sobre os alunos. Cada conjunto de dados é representado por um mapa. O mapa "studentData1" contém informações sobre as notas dos alunos para a primeira metade do semestre, enquanto o mapa "studentData2" contém informações sobre as notas para a segunda metade do semestre.
//
//Sua tarefa é criar uma função chamada "mergeStudentData" que recebe os mapas "studentData1" e "studentData2" como parâmetros e retorna um novo mapa que contém as informações combinadas dos dois conjuntos de dados.
//
//O objetivo é combinar as informações de cada aluno, preservando o nome e a idade, e atualizando as matérias e notas de acordo com o mapa mais recente. Lembre-se de que um aluno pode estar matriculado em diferentes matérias em cada metade do semestre.

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1 map[string]Student, studentData2 map[string]Student) map[string]Student {
	// Seu código aqui
	package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func mergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
	mergedData := make(map[string]Student)

	// Percorre os alunos no mapa studentData1
	for id, student := range studentData1 {
		mergedData[id] = student
	}

	// Percorre os alunos no mapa studentData2
	for id, student := range studentData2 {
		// Verifica se o aluno já existe no mapa mergedData
		if _, ok := mergedData[id]; ok {
			// O aluno já existe, atualiza as matérias e notas de acordo com o mapa mais recente
			for subject, grade := range student.Subjects {
				mergedData[id].Subjects[subject] = grade
			}
		} else {
			// O aluno não existe no mapa mergedData, adiciona-o diretamente
			mergedData[id] = student
		}
	}

	return mergedData
}

func readStudentData() map[string]Student {
	studentData := make(map[string]Student)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite o ID do aluno (ou 'q' para sair): ")
		idInput, _ := reader.ReadString('\n')
		idInput = strings.TrimSpace(idInput)

		if idInput == "q" {
			break
		}

		fmt.Print("Digite o nome do aluno: ")
		nameInput, _ := reader.ReadString('\n')
		nameInput = strings.TrimSpace(nameInput)

		fmt.Print("Digite a idade do aluno: ")
		ageInput, _ := reader.ReadString('\n')
		ageInput = strings.TrimSpace(ageInput)
		age, _ := strconv.Atoi(ageInput)

		student := Student{
			Name:     nameInput,
			Age:      age,
			Subjects: make(map[string]float64),
		}

		fmt.Println("Digite as notas das matérias (digite 'q' para sair)")

		for {
			fmt.Print("Matéria: ")
			subjectInput, _ := reader.ReadString('\n')
			subjectInput = strings.TrimSpace(subjectInput)

			if subjectInput == "q" {
				break
			}

			fmt.Print("Nota: ")
			gradeInput, _ := reader.ReadString('\n')
			gradeInput = strings.TrimSpace(gradeInput)
			grade, _ := strconv.ParseFloat(gradeInput, 64)

			student.Subjects[subjectInput] = grade
		}

		studentData[idInput] = student
	}

	return studentData
}

func main() {
	fmt.Println("### Coleta de dados para o primeiro conjunto de alunos ###")
	studentData1 := readStudentData()

	fmt.Println("\n### Coleta de dados para o segundo conjunto de alunos ###")
	studentData2 := readStudentData()

	// Mescla os dados dos alunos
	mergedData := mergeStudentData(studentData1, studentData2)

	// Imprime os dados mesclados
	for id, student := range mergedData {
		fmt.Printf("Aluno ID: %s\n", id)
		fmt.Printf("Nome: %s\n", student.Name)
		fmt.Printf("Idade: %d\n", student.Age)
		fmt.Println("Matérias e Notas:")
		for subject, grade := range student.Subjects {
			fmt.Printf("%s: %.1f\n", subject, grade)
		}
		fmt.Println()
	}
}

	return nil
}
