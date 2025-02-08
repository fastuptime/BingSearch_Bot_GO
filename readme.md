# 🔍 Bing Search Bot - GO

Bing Search Bot is a Go application that searches Bing for a specific keyword and returns the results in JSON format. This bot can search based on a specific language and region, and it paginates the results for easy navigation.

---

## 🚀 Features

- **🔑 Keyword Search**: Searches Bing for a specific keyword.
- **🌍 Language and Region Support**: Filters results based on a specific language and region.
- **📄 Pagination Support**: Paginates results and allows users to navigate through multiple pages.
- **📊 JSON Output**: Returns results in a structured JSON format for easy integration with other tools.
- **⚙️ Customizable**: Allows users to specify the number of results per page and the target language/region.

---

## 🛠️ Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/fastuptime/BingSearch_Bot_GO.git
   cd BingSearch_Bot_GO
   ```

2. **Install Dependencies**:
   Make sure you have Go installed. Then, install the required Go packages:
   ```bash
   go mod tidy
   ```

3. **Run the Bot**:
   ```bash
   go run main.go
   ```

---

## 🖥️ Usage

1. **Enter a Keyword**:
   When prompted, enter the keyword you want to search for.

2. **Specify Language and Region**:
   Enter the language code (e.g., `en` for English, `tr` for Turkish) and the region code (e.g., `US` for the United States, `TR` for Turkey).

3. **View Results**:
   The bot will display the search results in JSON format. You can navigate through multiple pages if available.

---

## 📝 Example

### Input:
```plaintext
Enter the keyword you want to search: OpenAI
Enter the language code (e.g., 'en' for English, 'tr' for Turkish): en
Enter the region code (e.g., 'US' for the United States, 'TR' for Turkey): US
```

### Output:
```json
{
  "results": [
    {
      "title": "OpenAI",
      "description": "OpenAI is an AI research and deployment company.",
      "url": "https://www.openai.com"
    },
    {
      "title": "OpenAI - Wikipedia",
      "description": "OpenAI is an artificial intelligence research laboratory.",
      "url": "https://en.wikipedia.org/wiki/OpenAI"
    }
  ],
  "current_page": 1,
  "next_page": 2
}
```

---

## 🌐 Supported Language and Region Codes

### Language Codes:
- `en`: English
- `tr`: Turkish
- `zh`: Chinese
- `es`: Spanish
- `fr`: French
- `de`: German
- `ru`: Russian
- `ja`: Japanese
- `ar`: Arabic

### Region Codes:
- `US`: United States
- `TR`: Turkey
- `CN`: China
- `DE`: Germany
- `FR`: France
- `JP`: Japan
- `RU`: Russia

---

## ⚠️ Notes

- Bing may sometimes ignore the language/region settings based on the user's IP address. If you encounter this issue, try using a VPN to access Bing from a different region.
- The bot includes a 5-second delay between requests to avoid overloading Bing's servers.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 💬 Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

---

## 📧 Contact

For questions or feedback, feel free to reach out:

- **GitHub**: [fastuptime](https://github.com/fastuptime)
