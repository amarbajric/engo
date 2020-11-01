#version 410
in vec3 vp;
out vec4 color;
void main() {
    color = vec4(clamp(vp, 0.0, 1.0), 1.0);
    gl_Position = vec4(vp, 1.0);
}