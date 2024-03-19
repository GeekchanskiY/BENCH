from dataclasses import dataclass

class Query:
    def __init__(self) -> None:
        pass

    def __repr__(self) -> str:
        pass

@dataclass(slots=True)
class Company:
    name: str
    rating: int | None
    description: str | None
    city: str | None
    domains: list[str] | None
    site_url: str | None


@dataclass(slots=True)
class ItemPreview:
    vacancy_name: str
    vacancy_link: str

    company_name: str
    company_link: str

    vacancy_city: str

    experience_min: int
    experience_max: int

    remote_job: bool

    contacts_preview: bool

