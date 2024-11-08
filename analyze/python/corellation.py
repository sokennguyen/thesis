import numpy as np
import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns
from sklearn.metrics import r2_score, root_mean_squared_error
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split

df = pd.read_csv('./firstVer.csv', delimiter=';')
print(df.head())

sns.heatmap(data=df.corr().round(2).abs(), annot=True)

