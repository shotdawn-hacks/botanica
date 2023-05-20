import pdfplumber
import json

with pdfplumber.open("test_files/2021.pdf") as pdf:
    result = {}
    temp = []
    first = True
    for i in range(18, 618):  # 617
        text = pdf.pages[i]

        clean_text = text.filter(
            lambda obj: obj["object_type"] == "char" and "Bold" in obj["fontname"]
        )

        for lines in clean_text.extract_text().split("\n"):
            if lines.isupper():
                if not first:
                    result[key] = temp
                    temp = []

                else:
                    first = False

                key = lines

            else:
                temp.append(lines)
    with open("test_files/bold.json", "w", encoding="utf-8") as f:
        json.dump(result, f, ensure_ascii=False)
        f.close()
