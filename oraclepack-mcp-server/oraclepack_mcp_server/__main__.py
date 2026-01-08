import argparse
import asyncio
from .server import mcp

def main():
    parser = argparse.ArgumentParser(description="Oraclepack MCP Server")
    parser.add_argument(
        "--transport", 
        choices=["stdio", "streamable-http"], 
        default="stdio",
        help="MCP transport to use (default: stdio)"
    )
    parser.add_argument(
        "--host", 
        default="localhost",
        help="Host to bind for streamable-http (default: localhost)"
    )
    parser.add_argument(
        "--port", 
        type=int, 
        default=8000,
        help="Port to bind for streamable-http (default: 8000)"
    )
    
    args = parser.parse_args()
    
    if args.transport == "stdio":
        mcp.run(transport="stdio")
    elif args.transport == "streamable-http":
        # FastMCP.run(transport="sse") is what maps to streamable-http in python SDK
        mcp.run(transport="sse", host=args.host, port=args.port)

if __name__ == "__main__":
    main()
