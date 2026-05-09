import yaml
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

with open("config.yaml", "r", encoding="utf-8") as f:
    config = yaml.safe_load(f)

# 加载参数（固化参数）
loading_args = {
    "pretrained_model_name_or_path": config["pretrained_model_name_or_path"],
    "load_in_4bit": config.get("load_in_4bit", False),
    "load_in_8bit": config.get("load_in_8bit", False),
    "device_map": config.get("device_map", "auto"),
    "low_cpu_mem_usage": config.get("low_cpu_mem_usage", True),
    "offload_folder": config.get("offload_folder", None),
    "torch_dtype": getattr(torch, config.get("torch_dtype", "float16")),
}
if loading_args["load_in_4bit"]:
    loading_args["bnb_4bit_compute_dtype"] = getattr(torch, config.get("bnb_4bit_compute_dtype", "float16"))
    loading_args["bnb_4bit_quant_type"] = config.get("bnb_4bit_quant_type", "nf4")
    loading_args["bnb_4bit_use_double_quant"] = config.get("bnb_4bit_use_double_quant", True)

model = AutoModelForCausalLM.from_pretrained(**loading_args)
tokenizer = AutoTokenizer.from_pretrained(config["pretrained_model_name_or_path"])
# 推理时的生成配置（可动态覆盖）
default_gen_config = config["generation"]

def generate_text(prompt, **overrides):
    gen_config = default_gen_config.copy()
    gen_config.update(overrides)
    # 修复：移除空的 stop_strings
    if "stop_strings" in gen_config and not gen_config["stop_strings"]:
        del gen_config["stop_strings"]
    inputs = tokenizer(prompt, return_tensors="pt").to(model.device)
    outputs = model.generate(**inputs, tokenizer=tokenizer, **gen_config)
    return tokenizer.decode(outputs[0], skip_special_tokens=True)
if __name__ == "__main__":
    prompt = "请写一个快速排序算法"
    result = generate_text(prompt)
    print("=" * 50)
    print(f"提示词: {prompt}")
    print("=" * 50)
    print(f"模型输出:\n{result}")
    print("=" * 50)
                 