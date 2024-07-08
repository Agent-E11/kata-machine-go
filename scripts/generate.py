#!/usr/bin/python3

"""
Generate one day of katas
"""

import shutil
from os import path
import glob

# Check if there is a src directory
if not path.exists("src"):
    print("error: could not find src directory")

day_num = 0

# Increment the day number until the path doesn't exist
while path.exists(f"day{day_num}"):
    day_num += 1

day_dir = f"day{day_num}"

print(f"creating katas in: {day_dir}")

# Copy all the files in src to the new day directory
shutil.copytree("src", day_dir)

for file in glob.glob(f"{day_dir}/**/*.go"):

    if path.isdir(file):
        continue

    with open(file, "r") as f:
        content = f.read()

    content = content.replace(
        "/kata-machine-go/src/",
        f"/kata-machine-go/{day_dir}/"
    )

    with open(file, "w") as f:
        f.write(content)
