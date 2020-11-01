#version 410
in vec3 vp;
out vec4 colour;
uniform mat4 transform;
void main() {
    colour = vec4(clamp(vp, 0.0, 1.0), 1.0);
    gl_Position = transform * vec4(vp, 1.0);
}