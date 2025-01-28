let output = document.getElementById("output");
let tab = document.getElementById("output_tab");
let input = document.getElementById("input")
let run_button = document.getElementById("run")
let example = document.getElementById("examples-select")
output.value = "";

document.addEventListener("keydown", (e) => {
	if (e.key == "Enter" && e.ctrlKey) {
		run_button.click()
	}
}) 

input.addEventListener('keydown', function(e) {
  if (e.key == 'Tab') {
    e.preventDefault();
    var start = this.selectionStart;
    var end = this.selectionEnd;

    this.value = this.value.substring(0, start) + "\t" + this.value.substring(end);

    this.selectionStart = this.selectionEnd = start + 1;
  }
});

example.onclick = () => {
	let name = example.value
	example.value = "def"
	loadExample(name)
}

function loadExample(val) {
	if (val == "def" || example_files == undefined) {
		return
	}
	let file = new FileReader()
	file.onloadend = (e) => {
		input.value = e.target.result
	}
	file.readAsText(example_files[val].file)
	
}

let example_files = undefined;
(async () => {
	let index = await fetch("examples/example_index.json")
	let all = await index.json()

	for (let i in all) {
		let res = await fetch(all[i].path)
		all[i].file = new Blob([await res.text()], {type: "text/plain"})

		let opt = document.createElement("option")
		opt.value = i
		opt.innerText = all[i].name
		example.appendChild(opt)
	}
	example_files = all

})();
