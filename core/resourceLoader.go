package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const SHADER_RESOURCES = "./res/shaders/"

func loadShader(fileName string) string {
	shaderProgram := fmt.Sprint()
	shaderFile, err := os.Open(SHADER_RESOURCES + fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while %s shader '%s': ", "loading", fileName), err)
	}

	defer func() {
		if err = shaderFile.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Error while %s shader file '%s': ", "closing", fileName), err)
		}
	}()

	bufferReader := bufio.NewReader(shaderFile)
	buffer := make([]byte, 1024)

	for {
		chunkLength, err := bufferReader.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(fmt.Sprintf("Error while %s shader file '%s': ", "reading", fileName), err)
			} else {
				shaderProgram += "\x00"
				break
			}
		}
		shaderProgram += string(buffer[0:chunkLength])
	}
	return shaderProgram
}