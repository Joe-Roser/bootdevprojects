#returns the number of words in the text
def num_words(text: str) -> int:
    return len(text.split())

#returns a dictionary that has all characters and their frequency
def char_stats(text: str) -> dict[str, int]:
    text = text.lower()
    chars = dict.fromkeys(set(text), 0)

    for char in list(text):
        chars[char] = chars[char] + 1

    return chars

#returns a list of dictionarys sorted max to min by char frequency
def organise_dist(d: dict[str, int]):
    dicts = []

    for kv in d.items():
        dicts.append(dict(char = kv[0], num = kv[1]))

    dicts.sort(reverse=True, key=lambda d : d["num"])
    return dicts
