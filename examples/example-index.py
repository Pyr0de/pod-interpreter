import os
import json

examples_dir = "examples"
out_file = os.path.join("bin", "example_index.json")

file_info = []

for i in os.listdir("examples"):
    if not i.endswith(".pod"):
        continue
    path = os.path.join("examples", i)
    file = open(path, "r")
    content = file.readline().strip()
    assert content.startswith("// ")

    name = content.split("// ")[1]
    file_info.append({
        "name": content.split("// ")[1],
        "path": path,
        })
    file.close()


if os.path.exists(out_file):
    f = open(out_file, "w")
else:
    f = open(out_file, "x")
f.write(json.dumps(file_info))
f.close()
