from langchain import hub
from langchain.agents import create_json_chat_agent, AgentExecutor
from langchain_openai import ChatOpenAI

def initialize_openai_agent_executor(tools):

    # Pull the prompt configuration
    prompt = hub.pull("hwchase17/react-chat-json")

    # Initialize the language model
    llm = ChatOpenAI(
            model="gpt-3.5-turbo-0125",
    )
        
    # Initialize the React Agent
    react_agent = create_json_chat_agent(llm, tools, prompt)
    
    # Create and return an AgentExecutor with the agent and tools
    agent_executor = AgentExecutor(
        agent=react_agent,
        tools=tools,
        verbose=True,
        max_iterations=5,
        handle_parsing_errors=True,
    )

    return agent_executor