package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func loadMesh(fileName string) Mesh {
	if !strings.HasSuffix(fileName, ".obj") {
		log.Fatalf("Error while loading mesh file: file format not supported for file '%s'", fileName)
	}
	vertices := make([]Vertex, 0)
	faces := make([]uint32, 0)

	meshObjFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while %s mesh obj file '%s': ", "loading", fileName), err)
	}

	defer func() {
		if err = meshObjFile.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Error while %s mesh obj file '%s': ", "closing", fileName), err)
		}
	}()

	regex, _ := regexp.Compile("([v|f])\\s+([\\d\\.E+-]+)\\s+([\\d\\.E+-]+)\\s+([\\d\\.E+-]+)")
	bufferReader := bufio.NewReader(meshObjFile)

	for {
		line, err := bufferReader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(fmt.Sprintf("Error while %s mesh obj file '%s': ", "reading", fileName), err)
			} else {
				break
			}
		} else {
			groups := regex.FindStringSubmatch(line)
			if len(groups) < 1 {
				continue
			}
			if groups[1] == "f" {
				for _, face := range groups[2:] {
					faces = append(faces, uint32(strToInt(face)) - 1)
				}
			} else if groups[1] == "v" {
				vertices = append(vertices, Vertex{Pos: Vector3f{
					X: strToFloat(groups[2]),
					Y: strToFloat(groups[3]),
					Z: strToFloat(groups[4]),
				}})
			}
		}
	}
	meshRes := Mesh{}
	meshRes.AddVertices(vertices, faces)
	return meshRes
}

func strToInt(numericVal string) int {
	val, err := strconv.Atoi(numericVal)
	if err != nil {
		log.Fatalf("Could not convert numeric string value to int: %s", err)
	}
	return val
}

func strToFloat(numericVal string) float64 {
	val, err := strconv.ParseFloat(numericVal, 64)
	if err != nil {
		log.Fatalf("Could not convert numeric string value to float: %s", err)
	}
	return val
}