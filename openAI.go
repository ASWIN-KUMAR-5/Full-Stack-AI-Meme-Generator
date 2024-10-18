from transformers import pipeline

# Load a pre-trained model for conversational AI
chatbot = pipeline('conversational', model='microsoft/DialoGPT-medium')

def chat():
    print("Chatbot: Hi! I'm your chatbot. How can I help you today?")
    while True:
        user_input = input("You: ")
        if user_input.lower() in ["exit", "quit"]:
            print("Chatbot: Goodbye!")
            break

        # Get a response from the model
        response = chatbot(user_input)
        print(f"Chatbot: {response[0]['generated_text']}")

if __name__ == "__main__":
    chat()
