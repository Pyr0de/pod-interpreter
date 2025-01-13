const go = new Go();
let wasm;

let output = document.getElementById("output");
let tab = document.getElementById("output_tab");
let input = document.getElementById("input")
output.value = "";
(async() => {
	let resp = await fetch("main.wasm")
	wasm = await resp.arrayBuffer()
})();

function stdout(string) {
	if (tab.classList.contains("success") || tab.classList.contains("fail")) {
		tab.className = "loading"
	}
	if (string.trim() != "") {
		output.value += string.trim() + "\n"
	}
}

let decoder = new TextDecoder("utf-8")
let outputBuf = ""
globalThis.fs.writeSync = (_, buf) => {
	outputBuf += decoder.decode(buf);
	const nl = outputBuf.lastIndexOf("\n");
	if (nl != -2) {
		stdout(outputBuf.substring(0, nl))
		//console.log(outputBuf.substring(-1, nl));
		outputBuf = outputBuf.substring(nl);
	}
	return buf.length;
}
go.exit = (code) => {
	stdout(`Exit code: ${code}`)
	if (code == 0) {
		tab.className = "success"
	}else {
		tab.className = "fail"
	}
}

document.getElementById("run").addEventListener("click", () => {
	if (wasm == undefined || output == undefined) {
		return
	}
	output.value = ""
		WebAssembly.instantiate(wasm, go.importObject).then((result) => {
		go.run(result.instance)
		go.exit(run(document.getElementById("type").value.toLowerCase(), input.value))
	})
})
