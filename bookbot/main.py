from stats import *
import sys

#returns the text from the file
def get_book_text(file_name: str) -> str: 
    with open(file_name) as f:
        return f.read()

# The main function innit
def main():
    text = get_book_text(sys.argv[1])

    print("""============ BOOKBOT ============
Analyzing book found at books/frankenstein.txt...
----------- Word Count ----------""")

    print(f"Found {num_words(text)} total words")


    print("--------- Character Count -------")

    chars = char_stats(text)
    chars = organise_dist(chars)
    for d in chars:
        if d["char"].isalpha():
            print(f"{d["char"]}: {d["num"]}")

    print("============= END ===============")

# call the main body
if len(sys.argv) == 1:
    print("Usage: python3 main.py <path_to_book>")
    sys.exit(1)
else:
    main()

