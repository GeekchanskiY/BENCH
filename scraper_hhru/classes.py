from dataclasses import dataclass, field
from config import SEARCH_STRING

from copy import deepcopy, copy

# @dataclass(slots=True)
class Query:

    available_query_args: frozenset = frozenset((
        'text', # str
        'page', # int
        'search_field', # values are in available_search_fields
        'enable_snippets', # false by default
        'part_time', # available_part_time
        'only_with_salary', # bool
        'salary', # ?int, 50K, 100K, 150K, 200K, 250K, custom
        'area', # int 
        'L_save_area', # bool
        'experience', # available_experience
        'schedule', # available_schedule
    ))

    available_search_fields: frozenset = frozenset((
        'name', 'company_name', 'description'
    ))

    available_part_time: frozenset = frozenset((
        'employment_part', # Not full day
        'start_after_sixteen', # Evening job
        'only_saturday_and_sunday', # Weekend job
        'employment_project', # Project job
        'from_four_to_six_hours_in_a_day', # 4-6 h/day
    ))

    available_experience: frozenset = frozenset((
        'between1And3',
        'between3And6',
        'moreThan6',
        'noExperience',
    ))

    available_schedule: frozenset = frozenset((
        'remote',
        'fullDay',
        'shift',
    ))


    def __init__(
            self,
            page: int|None = None,
            search_field: str|list|None = None,
            search_query: str|None = None,
            enable_snippets: bool = False,
            part_time: str|list|None = None,
            only_with_salary: bool = False,
            salary: int|None = None,
            area: int|None = None,
            experience: str|None = None,
            schedule: str|list|None = None
        ) -> None:
        '''
            Initialization & validation
        '''
        self.search_url = SEARCH_STRING


        self.area = area

        if experience is not None:
            if type(experience) is not str:
                raise AttributeError('Experience must be str!')
            if experience not in self.available_experience:
                raise AttributeError('Invalid Experience')
        self.experience = experience

        self.schedule = schedule
        self.salary = salary
        self.only_with_salary = only_with_salary
        self.part_time = part_time

        if page is not None and type(page) != int:
            raise AttributeError('Page must be int or None')
        self.page = page

        if type(search_query) is not str:
            raise AttributeError('Search query must be a str')
        self.text = search_query

        self.enable_snippets = enable_snippets
        
        if search_field is not None:
            
            if type(search_field) == str:
                if search_field not in self.available_search_fields:
                    raise AttributeError('Invalid search_field')
                self.search_fields = (search_field,)
            elif type(search_field) == list or type(search_field) == set:
                search_field_param: str
                for search_field_param in search_field:
                    if search_field_param not in self.available_search_fields:
                        raise AttributeError('Invalid search_field in iterable')
                del search_field_param

                self.search_fields = set(search_field)
                
            else:
                raise AttributeError('search_field must be None, str or list[str]')
        else:
            self.search_fields = set(self.available_search_fields)
        
        

    def fill_query_args(self):
        self.query_args: list = list()
        if self.area is not None:
            # L_save_args required if search by area enabled
            self.query_args.append(
                ('L_save_args', True)
            )
    
        if self.page is not None:
            self.page = self.page
            self.query_args.append(
                ('page', self.page)
            )
        
        if self.experience is not None:
            self.query_args.append(
                ('experience', self.experience)
            )
        
        if self.text is not None:
            
            self.query_args.append(
                ('text', self.text)
            )


    def get_next_page(self) -> 'Query':
        new_query: Query = Query(
            page=self.page + 1,
            search_field=self.search_fields,
            search_query=self.text,
            enable_snippets=self.enable_snippets,
            part_time=self.part_time,
            only_with_salary=self.only_with_salary,
            salary=self.salary,
            area=self.area,
            experience=self.experience,
            schedule=self.schedule
        )
        
        return new_query

    def __str__(self) -> str:
        self.fill_query_args()
        for q in self.query_args:
            
            key, value = q
            self.search_url += f'&{key}={value}'
        return self.search_url

    def __repr__(self) -> str:
        return str(self)


@dataclass(slots=True)
class Company:
    name: str
    rating: int | None
    description: str | None
    city: str | None
    domains: set[str] | None
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

if __name__ == '__main__':
    q = Query(
        page=4,
        search_query='python'
    )
    print(str(q))
    q1 = q.get_next_page()
    print(str(q1))