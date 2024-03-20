from dataclasses import dataclass, field
from config import SEARCH_STRING
from aiohttp import ClientSession
from datetime import datetime

from copy import deepcopy, copy

# @dataclass(slots=True)
class Query:
    '''
        HH.ru Query Builder
    '''

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

        if schedule is not None and schedule not in self.available_schedule:
            raise AttributeError('Incorrect schedule!')
        self.schedule = schedule

        self.salary = salary

        self.only_with_salary = only_with_salary

        if part_time is not None and part_time not in self.available_part_time:
            raise AttributeError('Incorrect part_time')
        self.part_time = part_time

        if page is not None and type(page) != int:
            raise AttributeError('Page must be int or None')
        if page is not None and page < 1:
            raise AttributeError('Page cant be < 1')
        self.page = page -1

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
                ('L_save_args', 'true')
            )
    
        if self.page is not None:
            self.query_args.append(
                ('page', str(self.page))
            )
        
        if self.experience is not None:
            self.query_args.append(
                ('experience', self.experience)
            )
        
        if self.text is not None:
            
            self.query_args.append(
                ('text', self.text)
            )
        
        if self.search_fields is not None:
            f: str
            for f in self.search_fields:
                self.query_args.append(
                    ('search_field', f)
                )
        
        if self.schedule is not None:
            self.query_args.append(
                ('schedule', self.schedule)
            )
        
        if self.salary is not None:
            self.query_args.append(
                ('salary', self.salary)
            )

        if self.only_with_salary == True:
            self.query_args.append(
                ('only_with_salary', 'true')
            )
        
        if self.part_time is not None:
            self.query_args.append(
                ('part_time', self.part_time)
            )
        
        if self.enable_snippets == True:
            self.query_args.append(
                'enable_snippets', 'true'
            )
        
        if self.area is not None:
            self.query_args.append(
                ('area', str(self.area))
            )

    def get_next_page(self) -> 'Query':
        self.page += 1

    def __str__(self) -> str:
        self.fill_query_args()
        self.search_url = SEARCH_STRING
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


@dataclass(slots=True, init=True)
class Vacancy:
    name: str
    company: str
    company_link: str
    link: str
    description: str
    skills: list[str]
    pub_date: str
    # adress: str
    
    # current_viewers: int
    required_experience_min: int
    


@dataclass(slots=True, init=True)
class ItemPreview:
    vacancy_name: str
    vacancy_link: str

    company_name: str
    company_link: str

    vacancy_city: str

    experience_min: int | None
    experience_max: int | None

    remote_job: bool | None

    contacts_preview: bool | None


class ItemHolder:
    def __init__(self) -> None:
        self.item_previews: list[ItemPreview] = list()
    
    def add_item_preview(self, new_item: ItemPreview) -> bool:
        exists: bool = False
        for item in self.item_previews:
            if item.vacancy_link == new_item.vacancy_link:
                exists = True
                break
        
        if not exists:
            self.item_previews.append(new_item)
        
        return exists
    
    def get_item_previews(self) -> list[ItemPreview]:
        return set(self.item_previews)

    def __len__(self) -> int:
        return len(self.item_previews)

itemHolder = ItemHolder()

def get_itemHolder():
    return itemHolder

if __name__ == '__main__':
    q = Query(
        page=4,
        search_query='python',
        salary=150_000,
        only_with_salary=True,
        experience='between1And3',
        area=113
    )
    print(str(q))
    q1 = q.get_next_page()
    print(str(q1))