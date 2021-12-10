import pandas as pd
df = pd.read_csv("NewsArticles_Top10Keywords.csv")
df.to_json("./news.json", orient="records")