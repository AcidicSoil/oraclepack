import os
from pathlib import Path
from typing import List, Union, Any
from pydantic import Field, field_validator
from pydantic_settings import BaseSettings, SettingsConfigDict

class Settings(BaseSettings):
    model_config = SettingsConfigDict(
        env_prefix="ORACLEPACK_",
        env_file=".env",
        extra="ignore"
    )

    bin: str = Field(default="oraclepack", description="Path to the oraclepack binary")
    # Use Union to prevent pydantic-settings from forcing JSON decode
    allowed_roots: Any = Field(
        default_factory=lambda: [Path.cwd()],
        description="List of allowed filesystem roots"
    )
    workdir: Path = Field(default_factory=Path.cwd, description="Working directory for execution")
    enable_exec: bool = Field(default=False, description="Enable execution tools")
    max_read_bytes: int = Field(default=65536, description="Max bytes for file read operations")
    character_limit: int = Field(default=32000, description="Max characters for tool outputs")

    @field_validator("allowed_roots", mode="before")
    @classmethod
    def parse_allowed_roots(cls, v: Any) -> List[Path]:
        if isinstance(v, str):
            # Support both colon and comma separation
            delimiter = ":" if ":" in v else ","
            return [Path(p.strip()) for p in v.split(delimiter) if p.strip()]
        if isinstance(v, list):
            return [Path(p) if isinstance(p, (str, Path)) else p for p in v]
        return v

settings = Settings()
