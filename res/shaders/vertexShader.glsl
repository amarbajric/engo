#version 410
in vec3 vp;
out vec4 colour;
uniform float uniformFloat;
void main() {
    colour = vec4(clamp(vp, 0.0, uniformFloat), 1.0);
    gl_Position = vec4(vp, 1.0);
}