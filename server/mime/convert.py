import json
import sys

# python3 convert.py src.json out.sql server.mime

if __name__ == "__main__":
    if len(sys.argv) < 4:
        print(
            "Usage: convert sourcefile outputfile database \n\n"
            "ERROR: missing arguments"
        )
        sys.exit(1)
    with open(sys.argv[1], "r") as f:
        data = json.loads(f.read())
        flipped = {}
        for m, l in data.items():
            for i in l:
                flipped[i.replace("*", ".*", 1) + "$"] = m
        with open(sys.argv[2], "w") as r:
            r.write(f"TRUNCATE {sys.argv[3]};\n")
            r.writelines([f"INSERT INTO {sys.argv[3]} (extension, mimetype) VALUES ('{ext}', '{typ}');\n" for ext, typ in flipped.items()])
            print("generated SQL")
