import re
import uuid
import json

class PDFExtractor:
    def __init__(self, base_directory):
        with open(base_directory+"2021.txt","r", encoding="utf-8") as f:
            self.text = [e.replace("\n", "") for e in f.readlines()]
            f.close()

        with open(base_directory+"name.txt","r", encoding="utf-8") as f:
            self.names = [e.replace("\n", "") for e in f.readlines()]
            f.close()

        with open(base_directory+"scientific_name.txt","r", encoding="utf-8") as f:
            self.scientific_names = [e.replace("\n", "") for e in f.readlines()]
            f.close()
            
        self.default_splitters = ["Описание.","Другие виды.","Ареал.","Экология.","Ресурсы.",
                                  "Возделывание.","Сырьё.","Химический состав.",
                                  "Фармакологические свойства.","Применение в медицине.","Литература"]
        self.init_indicies()
        self.init_text_between_two()
    
    def default_json(self):
        plants_json = {
        "_id": "",
        "name": "",
        "scientific_name":"",
        "family":"",
        "scientific_family":"",
        "habitats": [],
        "regions": [],
        "soils": [],
        "is_global_redbook": False,
        "sowing": {
            "start": "",
            "end": "",
            "conditions": "",
            "description": ""
        },
        "harvesting": {
            "start": "",
            "end": "",
            "conditions": "",
            "description": ""
        },
        "chemical_composition":[{
            "name":"",
            "value":""
        }],
        "medical_preparations": "",
        "annual_demand_tons": 0
    }
        return plants_json.copy()
    
    # Searching words indecies
    def init_indicies(self):
        name_ind = {}
        for name in self.names:
            is_first = True
            for i, line in enumerate(self.text):
                if name in line:
                    if is_first:
                        is_first = False
                        continue
                    name_ind[name] = i
        self.name_ind = name_ind
        
    def init_text_between_two(self):
        all_text_lists = {}
        temp_text = []
        for i in range(len(self.names)-1):
            temp_text = []
            for j, line in enumerate(self.text):
                if j > self.name_ind[self.names[i]] and j < self.name_ind[self.names[i+1]]:
                    temp_text.append(line)

            all_text_lists[self.names[i]] = temp_text

        self.all_text_lists = all_text_lists
    
    def list_to_json(self, lst):
        new_struct = self.default_json()
        new_struct["_id"] = str(uuid.uuid4())
        for ls in lst:
            s = ls.replace("\n", "")
            if ls.startswith("Семейство"):
                fams = re.split("\s*[—–]\s+", s)
                new_struct["family"] = fams[0][len("Семейство"):]
                try:
                    new_struct["scientific_family"] = fams[1]
                except:
                    print(fams)
                break

        return new_struct
    
    # Exporting json
    def export_to_mongo(self, path):
        all_text = {}
        for key in self.all_text_lists.keys():
            all_text[key] = self.list_to_json(self.all_text_lists[key])
            all_text[key]["name"] = key.lower()
            all_text[key]["scientific_name"] = self.scientific_names[self.names.index(key)]
        
        with open(path, "w", encoding="utf-8") as f:
            json.dump([value for value in all_text.values()], f, ensure_ascii=False)
            f.close()
            print("Done")
        self.all_text = all_text
            
    def parse_one_plant_text(self, name, splitters):
        d = {
            "_id":str(uuid.uuid4()),
            "name":name,
            "descr":[]
        }
        base_struct = {
            "field":"",
            "text":""
        }
        
        try:
            t = "".join(self.all_text_lists[name])
        except:
            return d
        
        t = t.replace("-\n", "")
        t = t.replace("\n", "")
        t = t.replace("  ", " ")
        t = t.replace("• ", " ")
        
        for splitter in splitters:
            if splitter not in t:
                splitters.remove(splitter)
        t_list = t.split(splitters[0])
        
        for splitter in splitters[1:]:
            new_list = []
            for t_e in t_list:
                new_list.extend(t_e.split(splitter))
            
            t_list = new_list
        
        for i, splitter in enumerate(splitters):
            bb = base_struct.copy()
            bb["field"] = splitter
            bb["text"] = t_list[i+1]
            
            d["descr"].append(bb)
            
        return d
    
    def parse_one_plant_another(self, name, splitters):

        d = {
            "_id":str(uuid.uuid4()),
            "name":name,
            "descr":[]
        }

        base_struct = {
            "field":"",
            "text":""
        }
        try:
            t = "".join(self.all_text_lists[name])
        except:
            return d
        t = t.replace("-\n", "")
        t = t.replace("\n", "")
        t = t.replace("  ", " ")
        t = t.replace("• ", " ")
        ind = {}
        
        for splitter in splitters:
            r = t.find(splitter)
            if r > 0:
                ind[splitter] = r
        ind = dict(sorted(ind.items(), key=lambda item: item[1]))
        
        keys = [key for key in ind.keys()]
        for i in range(len(keys)-1):
            bb = base_struct.copy()
            bb["field"] = keys[i]
            bb["text"] = t[ind[keys[i]]:ind[keys[i+1]]]
            
            d["descr"].append(bb)
        
        return d
            
    def export_all_plants_text(self, path):
        xprt = []
        for name in self.names:
            xprt.append(self.parse_one_plant_another(name, self.default_splitters))
            
        with open(path, "w", encoding="utf-8") as f:
            json.dump(xprt, f,ensure_ascii=False)
            f.close()
            print(f"Exported at {path}")
            

