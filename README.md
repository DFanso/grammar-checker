# Grammar Checker 🎯

A powerful and intelligent grammar checking tool built with Go and Google's Gemini AI, designed to help you write better and more accurate text. This tool provides detailed grammar analysis, corrections, and explanations in real-time.

## ✨ Features

- Real-time grammar and spelling checks using Gemini AI
- Detailed explanations for each correction
- Comprehensive grammar rules and their applications
- Color-coded output for better readability
- Interactive command-line interface
- Professional teaching tone with thorough feedback
- High performance with minimal resource usage

## 🚀 Installation

### Prerequisites

- Go 1.16 or higher
- Gemini API key (for AI-powered language processing)

### Steps

1. Clone the repository:
```bash
git clone https://github.com/dfanso/grammar-checker.git
cd grammar-checker
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your environment variables:
```bash
cp .env.example .env
# Add your Gemini API key to the .env file
```

## 💻 Usage

Run the grammar checker in interactive mode:

```bash
./grammar-checker
```

The tool will start in interactive mode where you can:
1. Type your text at the prompt
2. Get instant feedback with:
   - Original text
   - Corrected version
   - Detailed explanations
   - Applicable grammar rules
3. Press Ctrl+C to exit

## 🛠️ Configuration

Configure the application through the `.env` file:

```env
GEMINI_API_KEY=your_gemini_api_key_here
```

## 📦 Project Structure

```
grammar-checker/
├── main.go           # Main application code with Gemini AI integration
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
├── .env.example     # Example environment configuration
└── .env             # Your environment configuration (git-ignored)
```

## 🤝 Technical Details

The application uses:
- Google's Gemini AI for advanced language processing
- Color-coded terminal output for better UX
- System prompts for consistent AI responses
- Efficient error handling and user input processing

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the terms of the license included in the repository.

## ⭐ Support

If you find this project helpful, please give it a star on GitHub!

## 🔗 Links

- [Report Bug](https://github.com/dfanso/grammar-checker/issues)
- [Request Feature](https://github.com/dfanso/grammar-checker/issues)

---

Made with ❤️ using Go and Gemini AI 
