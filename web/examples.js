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
