import pdfplumber
import numpy as np
temp = []
result = {}

with pdfplumber.open("test_files/2021.pdf") as pdf:
    for i in range(18, 22): #617
            text = pdf.pages[i]  
            clean_text = text.filter(lambda obj: obj["object_type"] == "char" and "Bold" in obj["fontname"])
            for lines in clean_text.extract_text().split("\n"):
                  if lines.isupper() and len(temp)>0:
                       result[lines] = temp
                  else:
                        temp.append(lines)


print(result)