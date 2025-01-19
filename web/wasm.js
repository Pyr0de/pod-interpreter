const go = new Go();
let wasm;
(async() => {
	let resp = await fetch("main.wasm")
	wasm = await resp.arrayBuffer()
})();

function stdout(string) {
	if (tab.classList.contains("success") || tab.classList.contains("fail")) {
		tab.className = "loading"
	}
	if (string.trim() != "") {
		output.innerHTML += string.trim() + "<br>"
	}
}

function stderr(string) {
	if (tab.classList.contains("success") || tab.classList.contains("fail")) {
		tab.className = "loading"
	}
	if (string.trim() != "") {
		let span = document.createElement("span")
		span.innerText = string.trim() + "\n"
		span.className = "err"
		output.appendChild(span)
	}
}

// program exit
go.exit = (code) => {
	if (code == 0) {
		tab.className = "success"
	}else {
		stderr(`Exit code: ${code}`)
		tab.className = "fail"
	}
};

// stdout and stderr
let decoder = new TextDecoder("utf-8");
let outputBuf = "";
globalThis.fs.writeSync = (fd, buf) => {
	outputBuf += decoder.decode(buf);
	const nl = outputBuf.lastIndexOf("\n");
	if (nl != -2) {
		let out = fd == 2 ? stderr : stdout
		out(outputBuf.substring(0, nl))
		outputBuf = outputBuf.substring(nl);
	}
	return buf.length;
}

// run program
run_button.addEventListener("click", () => {
	if (wasm == undefined || output == undefined) {
		return
	}
	output.innerText = "";
	output.innerHTML = "";
	WebAssembly.instantiate(wasm, go.importObject).then((result) => {
		go.run(result.instance)
		go.exit(run(document.getElementById("type").value, input.value))
	})
})
