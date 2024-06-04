export function makeScriptsExecutable(el) {
    el.querySelectorAll("script").forEach(script => {
        const clone = document.createElement("script")

        for (const attr of script.attributes) {
            clone.setAttribute(attr.name, attr.value)
        }

        clone.text = script.innerHTML
        script.remove()
        document.head.appendChild(clone)
    })
}

export function render_app(path, arg0) {
    const temp_body = document.createElement('body');
    temp_body.innerHTML = go_wasm_handler(path, arg0);
    makeScriptsExecutable(temp_body)
    document.documentElement.replaceChild(temp_body, document.body);
}

export function render_test(args) {
    const { key, ...rest } = args
    render_app(key, JSON.stringify(rest))
}
