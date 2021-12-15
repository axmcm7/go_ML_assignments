*Assignment - Sales Forecast* 

The data scientists at KMart have collected 2013 sales data for 1559
products across 10 stores in different cities. Also, certain attributes
of each product and store have been defined.

Kmart wants to better understand the properties of products and stores
which play a key role in increasing sales.

_Goal_: The aim is to build a predictive model to find out the sales
of each product at a particular store.

Questions to answer:
====================
1.  What steps should be taken for pre-processing the data?

    In my `preprocess_data.py` script I first dropped the "Item_Identifier" column, because it immediately jumped out to me that this column would not be useful in predicting sales on the test set. I then filled missing values depending on the type of each column: if categorical, I filled them with the mode value. If numerical, I filled them with the mean value of the column.
    I noticed that "Item_Fat_Content" had inconsistent labeling (I checked the other columns as well), so I manually cleaned up the labels. I then one-hot encoded the categorical variables, then reordered the columns and dropped the "Outlet_Establishment_Year" column.

2.  Explain your modeling approach:

    a.  Choice of ML algorithm

    I decided to use basic linear regression because this appeared to be a fairly standard regression problem. I tried various Go library implementations of linear regression, finding each had its pros and cons. The [implementation]('https://github.com/sajari/regression') I ultimately used was fairly straightforward to use, with the caveat that it lacked advanced functionality such as being able to use l1 or l2 regularization to constrain the resultant weights.

    b.  Feature Selection

    I didn't have too much time to spend experimenting with feature selection, so I instead used my intuition in dropping several columns that logically should not be causally related to an item's sales. With more time, I could've played around with the [regression library's]('https://github.com/sajari/regression') feature crosses, which allow for more complex regression models.

3.  How would you evaluate the performance of your model? Explain the
    choice of metrics used.

    I evaluated the performance of my model on the R-squared value, which came out to be 0.568868, a decent value for a regression analysis. My coefficients turned out to be larger than I anticipated, which I suspect is partly due to my Go library's implementation of regression in conjunction with my preprocessed training data containing many one-hot encoded variables.

4.  Show plots for visualizing your model performance.

    I did not have the time to learn a Go visualization library for this part. In python, this would have been much more straightforward and easy, not least due to the multitude of great visualization libraries available.

Data:
=====
Please note that the data may have missing values as some stores might
not report all the data due to technical glitches. Hence, it will be
required to treat them accordingly.

We have train (8523) and test (5681) data set, train data set has both
input and output variable(s). You need to predict the sales for test
data set.

_Datafiles_: test\_kmart.csv, train\_kmart.csv

Item\_Identifier: Unique product ID

Item\_Weight: Weight of product

Item\_Fat\_Content: Whether the product is low fat or not

Item\_Visibility: The % of total display area of all products in a store
allocated to the particular product

Item\_Type: The category to which the product belongs

Item\_MRP: Maximum Retail Price (list price) of the product

Outlet\_Identifier: Unique store ID

Outlet\_Establishment\_Year: The year in which store was established

Outlet\_Size: The size of the store in terms of ground area covered

Outlet\_Location\_Type: The type of city in which the store is located

Outlet\_Type: Whether the outlet is just a grocery store or some sort of
supermarket

Item\_Outlet\_Sales: Sales of the product in the particular store. This
is the outcome variable to be predicted.
