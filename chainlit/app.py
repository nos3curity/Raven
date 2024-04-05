from langchain.agents import AgentExecutor
import chainlit as cl
import agent as custom_agents
from langchain.callbacks.base import BaseCallbackHandler
from langchain.schema.runnable.config import RunnableConfig

@cl.on_chat_start
def start():
    chat_profile = cl.user_session.get("chat_profile")

    tools = []

    agent_executor = custom_agents.initialize_openai_agent_executor(tools)

    # Store the AgentExecutor in the user session
    cl.user_session.set("agent", agent_executor)

@cl.on_message
async def main(message: cl.Message):

    # Retrieve the AgentExecutor from the user session
    agent = cl.user_session.get("agent")  # type: AgentExecutor
    
    # Asynchronously invoke the agent with the message content
    res = await agent.ainvoke(
        {"input": message.content},
        config=RunnableConfig(callbacks=[cl.LangchainCallbackHandler(stream_final_answer=True)]),
    )

    await cl.Message(content=res['output']).send()