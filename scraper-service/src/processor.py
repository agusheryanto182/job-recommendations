import pandas as pd
import numpy as np

class JobDataProcessor:
    def __init__(self):
        pass
        
    def clean_data(self, df):
        """Clean and process job data"""
        # Remove duplicates
        df = df.drop_duplicates()
        
        # Clean text fields
        text_columns = ['Title', 'Company', 'Location', 'Description']
        for col in text_columns:
            df[col] = df[col].str.strip()
            
        return df
        
    def analyze_data(self, df):
        """Analyze job data"""
        analysis = {
            'total_jobs': len(df),
            'companies': df['Company'].value_counts(),
            'locations': df['Location'].value_counts(),
            'levels': df['Level'].value_counts()
        }
        return analysis