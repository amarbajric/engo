package core

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
	"strings"
)


type Shader struct {
	program uint32
}

type Status struct {
	statusCode int32
	errorMessage error
}

func (s *Shader) Init() {
	s.program = gl.CreateProgram()

	if s.program == 0 {
		log.Fatal("Shader creation failed: Could not find valid memory location in init phase")
	}
}

func (s *Shader) Bind() {
	gl.UseProgram(s.program)
}

func (s *Shader) AddVertexShader(programSource string) {
	s.addProgram(programSource, gl.VERTEX_SHADER)
}

func (s *Shader) AddFragmentShader(programSource string) {
	s.addProgram(programSource, gl.FRAGMENT_SHADER)
}
func (s *Shader) AddGeometryShader(programSource string) {
	s.addProgram(programSource, gl.GEOMETRY_SHADER)
}

func (s *Shader) CompileShader() {
	gl.LinkProgram(s.program)

	if _, linkErr := getStatus(s.program, gl.LINK_STATUS); linkErr != nil {
		log.Fatal(linkErr)
	}

	/** Only for debugging purpose IF a VAO is bound beforehand, else it will throw an exception **/
	//gl.ValidateProgram(s.program)
	//if _, validateErr := getStatus(s.program, gl.VALIDATE_STATUS); validateErr != nil {
	//	log.Fatal(validateErr)
	//}
}

func (s *Shader) addProgram(programSource string, shaderType uint32) {
	shader := gl.CreateShader(shaderType)

	if shader == 0 {
		log.Fatal("shader creation failed: Could not find valid memory location while adding shader")
	}

	compiledSource, free := gl.Strs(programSource)
	gl.ShaderSource(shader, 1, compiledSource, nil)
	free()
	gl.CompileShader(shader)

	if _, err := getStatus(shader, gl.COMPILE_STATUS); err != nil {
		log.Fatal(err)
	}
	gl.AttachShader(s.program, shader)
}

func getShaderStatus(shader uint32, statusCheckType *uint32, status *Status) {
	gl.GetShaderiv(shader, *statusCheckType, &status.statusCode)
	if status.statusCode == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		logMessage := strings.Repeat("\x00", int(logLength) + 1)
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logMessage))

		status.errorMessage = fmt.Errorf("shader status error: %v", logMessage)
	}
}

func getProgramStatus(program uint32, statusCheckType *uint32, status *Status) {
	gl.GetProgramiv(program, *statusCheckType, &status.statusCode)
	if status.statusCode == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		logMessage := strings.Repeat("\x00", int(logLength) + 1)
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(logMessage))

		status.errorMessage = fmt.Errorf("program status error: %v", logMessage)
	}
}

func getStatus(ptr uint32, statusCheckType uint32) (int32, error) {
	status := Status {
		statusCode:   -1,
		errorMessage: nil,
	}

	if statusCheckType == gl.COMPILE_STATUS {
		getShaderStatus(ptr, &statusCheckType, &status)
	} else if statusCheckType == gl.LINK_STATUS || statusCheckType == gl.VALIDATE_STATUS {
		getProgramStatus(ptr, &statusCheckType, &status)
	} else {
		status.errorMessage = fmt.Errorf("unsupported opengl status type '%d'", statusCheckType)
	}
	return status.statusCode, status.errorMessage
}
