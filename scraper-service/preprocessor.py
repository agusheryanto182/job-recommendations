import re
from datetime import datetime, timedelta
from typing import Dict, List, Any, Optional
import json
import csv
from location_standardizer import LocationChecker

class JobDataPreprocessor:
    def __init__(self):
        """
        Initialize JobDataPreprocessor object

        This is a no-op, but it's included for consistency with other classes.
        """
        self.location_standardizer = LocationChecker()
        
    def _extract_keywords(self, text: str) -> List[str]:
        """Extract important keywords from text"""
        if not text:
            return []
            
        try:
            # Common tech keywords to look for
            tech_keywords = {
                # Programming Languages
                'python', 'java', 'javascript', 'typescript', 'php', 'ruby', 'go',
                'c++', 'c#', 'swift', 'kotlin', 'rust', 'scala', 'c', 'perl', 'haskell',
                'lua', 'dart', 'groovy', 'r', 'matlab', 'julia', 'cobol', 'fortran',
                'assembly', 'prolog', 'erlang', 'elixir', 'clojure', 'f#', 'vb.net', 'golang'
                
                # Web Technologies
                'html', 'css', 'sass', 'less', 'webpack', 'babel', 'npm', 'yarn',
                'pnpm', 'vite', 'rollup', 'parcel', 'grunt', 'gulp', 'browserify',
                'webassembly', 'pwa', 'web components', 'websocket', 'graphql', 'rest',
                
                # Frameworks & Libraries
                'react', 'angular', 'vue', 'django', 'flask', 'spring', 'laravel',
                'express', 'node.js', 'next.js', 'nuxt.js', 'flutter', 'rails',
                'svelte', 'remix', 'gatsby', 'fastapi', 'nest.js', 'strapi',
                'asp.net core', 'symfony', 'codeigniter', 'yii', 'cake php',
                'backbone.js', 'ember.js', 'meteor', 'jquery', 'bootstrap', 'tailwind',
                'material-ui', 'chakra-ui', 'ant design', 'redux', 'mobx', 'vuex', 'pinia',
                
                # Databases & Storage
                'sql', 'mysql', 'postgresql', 'mongodb', 'redis', 'elasticsearch',
                'oracle', 'sqlite', 'nosql', 'cassandra', 'couchdb', 'mariadb',
                'dynamodb', 'firebase', 'neo4j', 'influxdb', 'cockroachdb', 'supabase',
                'planetscale', 's3', 'minio', 'graphql', 'prisma', 'sequelize', 'typeorm',
                
                # Cloud & DevOps
                'aws', 'azure', 'gcp', 'docker', 'kubernetes', 'jenkins', 'ci/cd',
                'git', 'github', 'gitlab', 'bitbucket', 'terraform', 'ansible', 'puppet',
                'chef', 'prometheus', 'grafana', 'elk stack', 'nginx', 'apache',
                'cloudflare', 'vercel', 'netlify', 'heroku', 'digitalocean', 'vagrant',
                'github actions', 'travis ci', 'circle ci', 'argocd', 'helm',
                
                # AI/ML & Data Science
                'tensorflow', 'pytorch', 'keras', 'scikit-learn', 'pandas', 'numpy',
                'matplotlib', 'seaborn', 'jupyter', 'hadoop', 'spark', 'kafka',
                'airflow', 'dbt', 'tableau', 'power bi', 'opencv', 'nltk', 'spacy',
                
                # Testing & QA
                'jest', 'mocha', 'cypress', 'selenium', 'junit', 'pytest', 'phpunit',
                'karma', 'jasmine', 'cucumber', 'postman', 'swagger', 'testng',
                'robot framework', 'appium', 'jmeter', 'k6', 'playwright',
                
                # Mobile Development
                'android', 'ios', 'react native', 'flutter', 'xamarin', 'ionic',
                'swift ui', 'jetpack compose', 'kotlin multiplatform', 'capacitor',
                'cordova', 'objective-c', 'android studio', 'xcode',
                
                # Security
                'oauth', 'jwt', 'https', 'ssl/tls', 'encryption', 'authentication',
                'authorization', 'penetration testing', 'owasp', 'cybersecurity',
                'firewall', 'vpn', 'identity management', 'keycloak', 'auth0',
                
                # Methodologies & Concepts
                'agile', 'scrum', 'kanban', 'tdd', 'rest', 'api', 'microservices',
                'mvc', 'oop', 'functional programming', 'domain driven design',
                'clean architecture', 'solid principles', 'design patterns',
                'event sourcing', 'cqrs', 'serverless', 'jamstack', 'clean code',
                
                # Tools & Others
                'jira', 'confluence', 'slack', 'trello', 'postman', 'swagger',
                'linux', 'unix', 'windows', 'macos', 'vs code', 'intellij', 'eclipse',
                'sublime text', 'vim', 'docker desktop', 'powershell', 'bash',
                'zsh', 'tmux', 'homebrew', 'apt', 'yum', 'chocolatey',
                
                # Soft Skills & Business
                'leadership', 'teamwork', 'communication', 'problem solving',
                'analytical', 'project management', 'agile', 'scrum', 'time management',
                'critical thinking', 'collaboration', 'presentation', 'negotiation',
                'stakeholder management', 'business analysis', 'product management',
                'technical writing', 'mentoring', 'consulting', 'client relations'
            }
            
            # Job-specific keywords
            job_keywords = {
                # Job Levels & Positions
                'junior', 'senior', 'lead', 'manager', 'architect', 'consultant',
                'intern', 'trainee', 'entry level', 'mid level', 'principal', 'staff',
                'director', 'vp', 'cto', 'cio', 'head of', 'chief', 'specialist',
                'administrator', 'analyst', 'associate', 'coordinator', 'supervisor',
                
                # Employment Types
                'full-time', 'part-time', 'remote', 'onsite', 'hybrid',
                'contract', 'permanent', 'temporary', 'freelance', 'internship',
                'apprenticeship', 'seasonal', 'project-based', 'fixed-term',
                
                # Company Types
                'startup', 'enterprise', 'product', 'service', 'agency', 'consulting',
                'corporation', 'non-profit', 'government', 'public sector', 'private sector',
                'multinational', 'small business', 'medium enterprise', 'fortune 500',
                
                # Job Functions
                'development', 'engineering', 'testing', 'deployment', 'maintenance',
                'support', 'design', 'implementation', 'research', 'analysis',
                'architecture', 'operations', 'security', 'infrastructure', 'networking',
                'database', 'cloud', 'devops', 'sre', 'quality assurance', 'ui/ux',
                
                # Role Types
                'frontend', 'backend', 'full stack', 'mobile', 'web', 'desktop',
                'embedded', 'systems', 'data science', 'machine learning', 'ai',
                'blockchain', 'iot', 'game development', 'ar/vr', 'cloud native',
                
                # Responsibilities
                'coding', 'programming', 'debugging', 'optimization', 'documentation',
                'mentoring', 'training', 'planning', 'estimation', 'review',
                'monitoring', 'troubleshooting', 'integration', 'migration', 'scaling',
                
                # Industry Sectors
                'fintech', 'healthtech', 'edtech', 'e-commerce', 'social media',
                'gaming', 'telecommunications', 'cybersecurity', 'automotive',
                'aerospace', 'manufacturing', 'retail', 'logistics', 'media',
                
                # Work Environment
                'agile environment', 'fast-paced', 'collaborative', 'innovative',
                'cross-functional', 'international', 'multicultural', 'flexible',
                'dynamic', 'team-oriented', 'self-managed', 'deadline-driven',
                
                # Benefits & Perks
                'competitive salary', 'equity', 'stock options', 'health insurance',
                'retirement plan', 'paid time off', 'professional development',
                'training budget', 'flexible hours', 'work-life balance',
                
                # Required Qualities
                'problem solver', 'team player', 'self-motivated', 'detail-oriented',
                'analytical', 'creative', 'innovative', 'proactive', 'adaptable',
                'independent', 'organized', 'leadership', 'communication skills',
                
                # Experience Requirements
                'entry-level', 'mid-career', 'experienced', 'expert',
                '0-2 years', '2-5 years', '5-8 years', '8+ years', '10+ years',
                'proven track record', 'hands-on experience', 'background in',
                
                # Education
                'bachelor', 'master', 'phd', 'degree', 'certification',
                'computer science', 'software engineering', 'information technology',
                'bootcamp', 'self-taught', 'professional certification',
                
                # Location Types
                'office-based', 'work from home', 'flexible location',
                'relocation', 'travel required', 'multiple locations',
                'headquarters', 'regional office', 'global', 'local'
            }
            
            # Majors and Fields of Study
            majors = {
                # Computer & Software Engineering
                'computer science', 'information technology', 'software engineering',
                'computer engineering', 'informatics', 'information systems',
                'software development', 'programming', 'web development',
                'mobile development', 'game development', 'systems engineering',
                
                # Data & Analytics
                'data science', 'machine learning', 'artificial intelligence',
                'big data', 'data analytics', 'business analytics',
                'data engineering', 'statistical computing', 'computational science',
                'business intelligence', 'predictive analytics', 'data mining',
                
                # Specialized Computing
                'cloud computing', 'edge computing', 'distributed systems',
                'parallel computing', 'quantum computing', 'high performance computing',
                'grid computing', 'fog computing', 'serverless computing',
                
                # Networks & Security
                'networking', 'network engineering', 'telecommunications',
                'cybersecurity', 'information security', 'network security',
                'computer networks', 'wireless communications', 'network administration',
                'ethical hacking', 'digital forensics', 'cryptography',
                
                # Database & Information Management
                'database management', 'database administration', 'data warehousing',
                'information management', 'knowledge management', 'content management',
                'records management', 'enterprise data management', 'master data management',
                
                # Systems & Infrastructure
                'systems administration', 'infrastructure management', 'devops',
                'site reliability engineering', 'platform engineering', 'system integration',
                'enterprise architecture', 'technical architecture', 'solutions architecture',
                
                # Emerging Technologies
                'blockchain', 'internet of things', 'augmented reality',
                'virtual reality', 'mixed reality', 'robotics',
                'autonomous systems', 'embedded systems', '5g technologies',
                'edge computing', 'quantum technology', 'nanotechnology',
                
                # Digital Media & Design
                'digital media', 'multimedia', 'interactive media',
                'digital design', 'user interface design', 'user experience design',
                'web design', 'digital animation', 'game design',
                'computer graphics', '3d modeling', 'visual computing',
                
                # Business & Management Information Systems
                'management information systems', 'business informatics',
                'information resource management', 'it service management',
                'enterprise systems', 'business process management',
                'digital transformation', 'technology management',
                
                # Applied Computing
                'bioinformatics', 'computational biology', 'health informatics',
                'medical informatics', 'environmental informatics', 'geoinformatics',
                'computational physics', 'computational chemistry', 'computational linguistics',
                
                # Mathematics & Theoretical Computer Science
                'computational mathematics', 'mathematical computing',
                'theoretical computer science', 'algorithms', 'computational theory',
                'discrete mathematics', 'applied mathematics', 'operations research',
                
                # Software Quality & Testing
                'software quality assurance', 'quality engineering',
                'test automation', 'performance engineering',
                'reliability engineering', 'software verification',
                
                # Project & Product Management
                'software project management', 'product management',
                'agile management', 'digital product management',
                'it project management', 'technical program management',
                
                # Related Engineering Fields
                'electronic engineering', 'electrical engineering',
                'mechatronics', 'automation engineering',
                'control systems', 'industrial automation',
                
                # Interdisciplinary Fields
                'cognitive science', 'human-computer interaction',
                'computational social science', 'digital humanities',
                'educational technology', 'information science',
                'knowledge engineering', 'systems science'
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
            
            # Add majors and fields of study found in text
            found_keywords.update(
                keyword for keyword in majors 
                if keyword in text_lower
            )
            
            # Extract additional keywords using regex patterns
            # Look for words that might be important but not in our predefined sets
            # additional_patterns = [
            #     # Camel case words (likely technical terms)
            #     r'[a-z]+[A-Z][a-zA-Z]*',
            #     # Words with dots (e.g., package names)
            #     r'\b[\w-]+\.[\w-]+(?:\.[\w-]+)*\b',
            #     # Words with version numbers
            #     r'\b[\w-]+\s*\d+(?:\.\d+)*\b',
            #     # Capitalized words (likely proper nouns)
            #     r'\b[A-Z][a-zA-Z]+\b'
            # ]
            
            # for pattern in additional_patterns:
            #     matches = re.findall(pattern, text)
            #     found_keywords.update(match.lower() for match in matches)
            
            # # Remove very short keywords and common words
            # common_words = {'the', 'and', 'or', 'in', 'at', 'to', 'for', 'of', 'with'}
            # keywords = {k for k in found_keywords if len(k) > 2 and k not in common_words}
            
            return sorted(list(found_keywords))
            
        except Exception as e:
            print(f"Warning: Error extracting keywords: {str(e)}")
            return []
        
    def _standardize_detail_job(self, text:str) -> str:
        """_summary_
        Standardize job details by handling various forms of missing/invalid value
        
        Args:
            text (str): input text to standardize

        Returns:
            None if text is empty/invalid, otherwise standardized text
        
        Examples:
            'Not Applicable' -> None
            'N/A' -> None
            'Not Available' -> None
            '' -> None
            'Entry Level' -> 'entry level'
        """
        
        try: 
            # Handle None or empty string
            if not text or not isinstance(text, str):
                return None
            
            # Standardize text
            text = text.lower().strip()
            
            invalid_values = [
                'not applicable', 'n/a', 'not available', '', 'n/a', '-', 'not specified', 'none'
            ]
            
            # check if text is one of the invalid values
            if text in invalid_values or text == '':
                return None
            
            return text
        except Exception as e:
            print(f"Warning: Error standardizing job details: {str(e)}")
            return None
            
    
    def _standardize_location(self, text:str) -> Dict:
        """_summary_
        Standardize location format with None for missing values
        """
        default_result ={
            'city': None,
            'province': None,
            'country': None
        }
        try:
            if not text or not isinstance(text, str):
                return default_result
            
            parts = [part.strip() for part in text.split(',')]
            
            result = {
                'city' : parts[0] if len(parts) > 0 and parts[0] else None,
                'province' : parts[1] if len(parts) > 1 and parts[1] else None,
                'country' : parts[2] if len(parts) > 2 and parts[2] else None,
            }
            
            return result
        except Exception as e:
            print(f"Warning: Error standardizing location: {str(e)}")
            return default_result
        
    def preprocess_job(self, job: Dict) -> Dict:
        """Preprocess job data into a better structure"""
        try:
            processed_job = {
                'posted_date': self._convert_date(job.get('posted_date')),
                # 'location': self.location_standardizer.standardize_location(job.get('location', '')),
                'seniority_level': self._standardize_detail_job(job.get('seniority_level', '')),
                'employment_level': self._standardize_detail_job(job.get('employment_level', '')),
                'job_function': self._standardize_detail_job(job.get('job_function', '')),
                'industries': self._standardize_detail_job(job.get('industries', '')),
                
                'processed_text': {
                    'keywords': self._extract_keywords(job.get('description', '')),
                    'description': job.get('description', ''),
                }
            }
            
            tempLocation = job.get('location')
            if tempLocation.split(',').__len__() > 2:
                processed_job['location'] = self._standardize_location(tempLocation)
            elif tempLocation.split(',').__len__() == 2:
                locations = tempLocation.split(',')
                locationType1 = self.location_standardizer.check_location_type(locations[0])
                locationType2 = self.location_standardizer.check_location_type(locations[1])
                if locationType1 == 'city' and locationType2 == 'country':
                    resProvince = self.location_standardizer.
                    processed_job['location'] = {
                        'city': locations[0],
                        'country': locations[1]
                    }
            
            return processed_job
        except Exception as e:
            print(f"Error preprocessing job: {str(e)}")
            # Return minimal job data instead of raw job
            return {
                'posted_date': job.get('posted_date', ''),
                'processed_text': job.get('processed_text', {}),
            }
            
    def _convert_date(self, text: str) -> str:
        """ Convert date string to ISO format
        Ex: '1 year' -> '2023-01-01'
        """
        
        if not text:
            return None
        
        result = re.search(r'(\d+) (year|month|week|day)s?', text)
        if result:
            value = int(result.group(1))
            unit = result.group(2)
            if unit == 'year':
                # should be like this 2023-01-01
                return (datetime.now() - timedelta(days=365 * value)).strftime('%Y-%m-%d')
            elif unit == 'month':
                return (datetime.now() - timedelta(days=30 * value)).strftime('%Y-%m-%d')
            elif unit == 'week':
                return (datetime.now() - timedelta(days=7 * value)).strftime('%Y-%m-%d')
            elif unit == 'day':
                return (datetime.now() - timedelta(days=value)).strftime('%Y-%m-%d')
        
        return None
    
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