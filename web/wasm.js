const go = new Go();
let wasm;
(async() => {
	let resp = await fetch("main.wasm")
	wasm = await resp.arrayBuffer()
})();

function stdout(string) {
	string = string.replace("\n", "<br>")
	if (string != "") {
		output.innerHTML += string + "<br>"
	}
}

function stderr(string) {
	if (tab.classList.contains("success") || tab.classList.contains("fail")) {
		tab.className = "loading"
	}
	if (string != "") {
		let span = document.createElement("span")
		span.innerText = string + "\n"
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
		out(outputBuf.substring(0, nl).trim())
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
	tab.className = "loading"
	WebAssembly.instantiate(wasm, go.importObject).then((result) => {
		go.run(result.instance)
		let start = new Date().getTime()
		go.exit(run(document.getElementById("type").value, input.value))
		let end = new Date().getTime()
		stdout("\nProgram exited...");
		stdout(`Time: ${end - start} ms`)
	})
})
