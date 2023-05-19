import pdfplumber
import numpy as np
demo = []
with pdfplumber.open("test_files/2021.pdf") as pdf:
    for i in range(20, 24):
        try:
            text = pdf.pages[i]  
            clean_text = text.filter(lambda obj: obj["object_type"] == "char" and "Bold" in obj["fontname"])
            print(clean_text.extract_text())
        except IndexError:
            print("")
            break
