from . import extractor

base_dir = "DIR_WITH_FORMATTING_FILES"

pdf_extractor = extractor.PDFExtractor(base_dir)
pdf_extractor.export_to_mongo("/home/hexzedels/Work/Hacks/botanica/test_files/mongo.json")
pdf_extractor.export_all_plants_text("/home/hexzedels/Work/Hacks/botanica/test_files/mongo_texts.json")