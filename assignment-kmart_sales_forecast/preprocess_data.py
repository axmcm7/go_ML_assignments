import pandas as pd
from sklearn.preprocessing import LabelBinarizer

train_df = pd.read_csv("./train_kmart.csv")
test_df = pd.read_csv("./test_kmart.csv")

# drop item_id column
train_df = train_df.drop(labels=["Item_Identifier"], axis=1)
test_df = test_df.drop(labels=["Item_Identifier"], axis=1)

cat_cols = ["Item_Fat_Content", "Item_Type", "Outlet_Identifier", "Outlet_Size", "Outlet_Location_Type", "Outlet_Type"]
num_cols = ["Item_Weight", "Item_Visibility", "Item_MRP", "Outlet_Establishment_Year", "Item_Outlet_Sales"]

# fill NaN values
# for categoricals: fill with mode value of column
# for numericals: fill with mean value of column
for col in cat_cols:
    train_df[col].fillna(train_df[col].mode()[0], inplace=True)
    train_df[col] = train_df[col].astype(str)
    test_df[col].fillna(test_df[col].mode()[0], inplace=True)
    test_df[col] = test_df[col].astype(str)

for col in num_cols:
    train_df[col].fillna(train_df[col].mean(), inplace=True)
    if col != "Item_Outlet_Sales":
        test_df[col].fillna(test_df[col].mean(), inplace=True)

# clean up categorical variable values for "Item_Fat_Content"
train_df = train_df.replace(
    {"Item_Fat_Content": {"LF": "Low Fat", "low fat": "Low Fat", "reg": "Regular"}})
test_df = test_df.replace(
    {"Item_Fat_Content": {"LF": "Low Fat", "low fat": "Low Fat", "reg": "Regular"}})

def convert_to_categoricals(df, col):
    x_pandas = pd.get_dummies(df[col])
    rename_col = lambda name: f"{col}-{name.replace(' ', '_')}"
    
    x_pandas.rename(columns=rename_col, inplace=True)

    df.drop(labels=[col], axis=1, inplace=True)
    df = pd.concat([df, x_pandas], axis=1)
    return df

# convert categorical variables to dummy variables
for col in cat_cols:
    train_df = convert_to_categoricals(train_df, col)
    test_df = convert_to_categoricals(test_df, col)

# place observed (dependent) variable as the last col
train_df["Sales"] = train_df["Item_Outlet_Sales"]
train_df.drop("Item_Outlet_Sales", axis=1, inplace=True)

# feature selection--columns to drop:
train_df.drop(["Outlet_Establishment_Year"], axis=1, inplace=True)
test_df.drop(["Outlet_Establishment_Year"], axis=1, inplace=True)


# normalize the data using mean normalization
#train_df = (train_df - train_df.mean()) / train_df.std()
#test_df = (test_df - test_df.mean()) / test_df.std()

# normalize the data using min-max normalization
#train_df = (train_df - train_df.min()) / (train_df.max() - train_df.min()) 
#test_df = (test_df - test_df.min()) / (test_df.max() - test_df.min())

train_df.to_json("./train.json", orient="records")
test_df.to_json("./test.json", orient="records")
#train_df.to_csv("./train.csv")
#test_df.to_csv("./test.csv")