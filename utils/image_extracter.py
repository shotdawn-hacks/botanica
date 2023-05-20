import pdfplumber
import re
import uuid
import json

export = []

f = int(input())
t = int(input())
save = input()

pdf_file_path = "test_files/2021.pdf"
save_file = open("test_files/images/"+save+".json", "w", encoding="utf-8")

with pdfplumber.open(pdf_file_path) as pdf:
    try:
        for i in range(f, t): #646
            print(i)
            page = pdf.pages[i]  
            page_images = page.images
            page_height = page.height
            all_names = []
            text = page.extract_text()

            for e, line in enumerate(text.split("\n")):
                names = re.findall("[А-Я][^A-Я]+", line)
                all_names.extend(names)

            # print(all_names)
            
            for j, image in enumerate(page_images):
                image_bbox = (image['x0'], page_height - image['y1'], image['x1'], page_height - image['y0'])
                cropped_page = page.crop(image_bbox)
                image_obj = cropped_page.to_image(resolution=400)
                u_uuid = str(uuid.uuid4())
                temp1 = {}
                temp1["_id"] = u_uuid
                temp1["name"] = all_names[11-j]
                export.append(temp1)

                image_obj.save("test_files/images/"+u_uuid+".jpg")
    except:
        json.dump(export, save_file, ensure_ascii=False)
    json.dump(export, save_file, ensure_ascii=False)