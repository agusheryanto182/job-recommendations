import re
from datetime import datetime
from typing import Dict, List, Any, Optional
import json
import csv
import uuid

class JobDataPreprocessor:
    """Class untuk preprocessing data pekerjaan dengan struktur yang lebih baik"""
    
    def __init__(self):
        # Initialize skill dictionaries
        self.skills_dict = {
            'skills': {
                'python', 'java', 'javascript', 'typescript', 'php', 'go', 'rust',
                'c++', 'c#', 'ruby', 'swift', 'kotlin', 'scala', 'perl', 'r',
                'dart', 'lua', 'haskell', 'erlang', 'julia',                 'react', 'angular', 'vue', 'django', 'flask', 'spring', 
                'laravel', 'express', 'next.js', 'nuxt.js', 'flutter',
                'tensorflow', 'pytorch', 'keras', 'symfony', 'rails',
                'svelte', 'gatsby', 'fastapi', 'nest.js','git', 'docker', 'kubernetes', 'jenkins', 'jira', 'aws', 
                'azure', 'gcp', 'terraform', 'ansible', 'nginx', 'linux',
                'webpack', 'babel', 'vite', 'ci/cd', 'github actions',
                'gitlab', 'bitbucket', 'prometheus', 'grafana', 'elk'
                'mysql', 'postgresql', 'mongodb', 'redis', 'elasticsearch',
                'cassandra', 'oracle', 'sql server', 'sqlite', 'dynamodb',
                'mariadb', 'neo4j', 'graphql', 'couchdb', 'firebase'
            }
        }
        
    
    def _extract_keywords(self, text: str) -> List[str]:
        """Extract important keywords from text"""
        if not text:
            return []
            
        try:
            # Common tech keywords to look for
            tech_keywords = {
                # Programming Languages
                'python', 'java', 'javascript', 'typescript', 'php', 'ruby', 'go',
                'c++', 'c#', 'swift', 'kotlin', 'rust', 'scala',
                
                # Web Technologies
                'html', 'css', 'sass', 'less', 'webpack', 'babel', 'npm', 'yarn',
                
                # Frameworks & Libraries
                'react', 'angular', 'vue', 'django', 'flask', 'spring', 'laravel',
                'express', 'node.js', 'next.js', 'nuxt.js', 'flutter', 'rails',
                
                # Databases
                'sql', 'mysql', 'postgresql', 'mongodb', 'redis', 'elasticsearch',
                'oracle', 'sqlite', 'nosql',
                
                # Cloud & DevOps
                'aws', 'azure', 'gcp', 'docker', 'kubernetes', 'jenkins', 'ci/cd',
                'git', 'github', 'gitlab', 'bitbucket',
                
                # Methodologies & Concepts
                'agile', 'scrum', 'kanban', 'tdd', 'rest', 'api', 'microservices',
                'mvc', 'oop', 'functional programming',
                
                # Tools & Others
                'jira', 'confluence', 'slack', 'trello', 'postman', 'swagger',
                'linux', 'unix', 'windows', 'macos',
                
                # Soft Skills & Business
                'leadership', 'teamwork', 'communication', 'problem solving',
                'analytical', 'project management', 'agile', 'scrum'
            }
            
            # Job-specific keywords
            job_keywords = {
                'junior', 'senior', 'lead', 'manager', 'architect', 'consultant',
                'full-time', 'part-time', 'remote', 'onsite', 'hybrid',
                'startup', 'enterprise', 'product', 'service',
                'development', 'engineering', 'testing', 'deployment',
                'maintenance', 'support', 'design', 'implementation'
            }
            
            # Convert text to lowercase for matching
            text_lower = text.lower()
            
            # Find all matching keywords
            found_keywords = set()
            
            # Add tech keywords found in text
            found_keywords.update(
                keyword for keyword in tech_keywords 
                if keyword in text_lower
            )
            
            # Add job keywords found in text
            found_keywords.update(
                keyword for keyword in job_keywords 
                if keyword in text_lower
            )
            
            # Extract additional keywords using regex patterns
            # Look for words that might be important but not in our predefined sets
            additional_patterns = [
                # Camel case words (likely technical terms)
                r'[a-z]+[A-Z][a-zA-Z]*',
                # Words with dots (e.g., package names)
                r'\b[\w-]+\.[\w-]+(?:\.[\w-]+)*\b',
                # Words with version numbers
                r'\b[\w-]+\s*\d+(?:\.\d+)*\b',
                # Capitalized words (likely proper nouns)
                r'\b[A-Z][a-zA-Z]+\b'
            ]
            
            for pattern in additional_patterns:
                matches = re.findall(pattern, text)
                found_keywords.update(match.lower() for match in matches)
            
            # Remove very short keywords and common words
            common_words = {'the', 'and', 'or', 'in', 'at', 'to', 'for', 'of', 'with'}
            keywords = {k for k in found_keywords if len(k) > 2 and k not in common_words}
            
            return sorted(list(keywords))
            
        except Exception as e:
            print(f"Warning: Error extracting keywords: {str(e)}")
            return []
    
    def preprocess_job(self, job: Dict) -> Dict:
        """Preprocess job data into a better structure"""
        try:
            processed_job = {
                # Basic Info
                'id': self._generate_job_id(),
                'title': self._clean_text(job.get('title', '')),
                'company': self._clean_text(job.get('company', '')),
                'link': job.get('link', ''),
                'posted_date': self._standardize_date(job.get('posted_date')),
                
                # Location
                'location': self._process_location(job.get('location', '')),
                
                'is_remote': job.get('is_remote', False),
                
                # Requirements
                'requirements': {
                    'education': self._extract_education(job.get('description', '')),
                    'experience_years': self._extract_experience_years(job.get('description', '')),
                },
                
                # Skills
                'skills': self._extract_all_skills(job.get('description', '')),

                # Processed Text
                'processed_text': {
                    'keywords': self._extract_keywords(job.get('description', '')),
                    'clean_description': self._clean_text(job.get('description', '')),
                }
            }
            return processed_job
        except Exception as e:
            print(f"Error preprocessing job: {str(e)}")
            # Return minimal job data instead of raw job
            return {
                'id': self._generate_job_id(),
                'title': job.get('title', ''),
                'company': job.get('company', ''),
                'link': job.get('link', ''),
                'description': job.get('description', ''),
                'location': job.get('location', ''),
                'posted_date': job.get('posted_date', '')
            }
    
    def _generate_job_id(self) -> str:
        """Generate unique job ID"""
        return str(uuid.uuid4())
    
    def _clean_text(self, text: str) -> str:
        """Clean text from special characters and extra spaces"""
        if not text:
            return ""
        text = re.sub(r'<[^>]+>', '', text)  # Remove HTML tags
        text = re.sub(r'\s+', ' ', text)      # Remove extra spaces
        return text.strip()
    
    def _standardize_date(self, date_str: Optional[str]) -> str:
        """Standardize date format to ISO"""
        if not date_str:
            return ""
        try:
            date_obj = datetime.strptime(date_str, '%Y-%m-%d')
            return date_obj.isoformat()
        except:
            return date_str
    
    def _process_location(self, location: str) -> Dict[str, str]:
        """Process location information"""
        location_info = {
            'address': location,
            'is_remote': self._detect_remote_work(location)
        }
        
        return location_info
    
    def _detect_remote_work(self, text: str) -> bool:
        """Detect if job is remote"""
        remote_indicators = {'remote', 'work from home', 'wfh', 'remote-first', 'fully remote'}
        return any(indicator in text.lower() for indicator in remote_indicators)
    
    
    def _convert_to_number(self, value: str) -> float:
        """Convert string number to float, handling different formats"""
        try:
            # Remove currency symbols and separators
            clean_value = re.sub(r'[^\d.]', '', value)
            return float(clean_value)
        except:
            return None
    
    def _detect_experience_level(self, text: str) -> str:
        """Detect experience level from job description"""
        text_lower = text.lower()
        
        patterns = {
            'entry': [
                r'0-2 years?', r'fresh graduate', r'entry level',
                r'junior', r'pemula', r'fresh', r'graduate'
            ],
            'mid': [
                r'2-5 years?', r'3-5 years?', r'mid level',
                r'intermediate', r'middle', r'experienced'
            ],
            'senior': [
                r'5\+? years?', r'senior', r'lead', r'manager',
                r'expert', r'principal', r'architect'
            ]
        }
        
        for level, pattern_list in patterns.items():
            if any(re.search(pattern, text_lower) for pattern in pattern_list):
                return level
                
        return 'not_specified'
    
    def _extract_all_skills(self, text: str) -> Dict[str, List[str]]:
        """Extract all types of skills from text"""
        text_lower = text.lower()
        skills = {
            'skills': [],
        }
        
        for category, skill_set in self.skills_dict.items():
            skills[category] = [
                skill for skill in skill_set 
                if skill in text_lower
            ]
            
        return skills
    
    def _extract_soft_skills(self, text: str) -> List[str]:
        """Extract soft skills from text"""
        soft_skills = {
            'communication', 'leadership', 'teamwork', 'problem solving',
            'analytical', 'creative', 'initiative', 'detail oriented',
            'time management', 'adaptability', 'collaboration'
        }
        
        text_lower = text.lower()
        return [skill for skill in soft_skills if skill in text_lower]
    
    def _extract_education(self, text: str) -> List[str]:
        """Extract education requirements"""
        education_patterns = [
            r"bachelor'?s?\s+degree",
            r"master'?s?\s+degree",
            r"ph\.?d",
            r"s1",
            r"s2",
            r"diploma",
            r"sarjana",
            r"magister"
            r"sma",
            r"smk"
        ]
        
        found_education = []
        text_lower = text.lower()
        
        for pattern in education_patterns:
            if re.search(pattern, text_lower):
                found_education.append(pattern.replace(r'\.?', '.'))
            else:
                found_education.append("sma")
        return found_education
    
    def _extract_experience_years(self, text: str) -> Dict:
        """Extract years of experience requirement"""
        experience = {
            'min': None,
        }
        
        # Pattern untuk mencari requirement pengalaman
        patterns = [
            r'(\d+)(?:-(\d+))?\s*(?:years?|tahun)',
            r'minimal\s*(\d+)\s*(?:years?|tahun)',
            r'at\s*least\s*(\d+)\s*(?:years?|tahun)'
        ]
        
        text_lower = text.lower()
        
        for pattern in patterns:
            match = re.search(pattern, text_lower)
            if match:
                experience['min'] = int(match.group(1))
            else:
                experience['min'] = 0
        return experience
    
    def _extract_languages(self, text: str) -> List[str]:
        """Extract required languages"""
        languages = {
            'english', 'indonesian', 'mandarin', 'japanese',
            'korean', 'french', 'german', 'spanish'
        }
        
        text_lower = text.lower()
        return [lang for lang in languages if lang in text_lower]
    
    
def save_to_json(jobs: List[Dict], filename: str):
    """Save jobs data to JSON file"""
    try:
        with open(filename, 'w', encoding='utf-8') as f:
            json.dump(jobs, f, ensure_ascii=False, indent=2)
        print(f"Successfully saved {len(jobs)} jobs to {filename}")
    except Exception as e:
        print(f"Error saving to JSON: {str(e)}")

def save_to_csv(jobs: List[Dict], filename: str):
    """Save jobs data to CSV file with flattened structure"""
    if not jobs:
        print("No jobs to save")
        return
        
    # Flatten nested dictionary
    flattened_jobs = []
    for job in jobs:
        flat_job = {}
        
        def flatten_dict(d, parent_key=''):
            for k, v in d.items():
                new_key = f"{parent_key}_{k}" if parent_key else k
                
                if isinstance(v, dict):
                    flatten_dict(v, new_key)
                elif isinstance(v, list):
                    flat_job[new_key] = '|'.join(str(x) for x in v)
                else:
                    flat_job[new_key] = v
                    
        flatten_dict(job)
        flattened_jobs.append(flat_job)
    
    # Write to CSV
    with open(filename, 'w', encoding='utf-8', newline='') as f:
        writer = csv.DictWriter(f, fieldnames=flattened_jobs[0].keys())
        writer.writeheader()
        writer.writerows(flattened_jobs)