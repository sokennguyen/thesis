import numpy as np
import matplotlib.pyplot as plt
from sklearn import linear_model

data = np.genfromtxt('data.csv', delimiter=',')
xp = data[:, 0]
yp = data[:, 1]
xp = xp.reshape(-1, 1)
yp = yp.reshape(-1, 1)

# regression
regr = linear_model.LinearRegression()
regr.fit(xp, yp)  # fitting the model = training the model = make a regression out of the x,y arrays
print(regr.coef_, regr.intercept_)

xval = np.full((1, 1), 0.5)  # (1,1) is the shape of the array, 0.5 is the value that will get filled into the array
yval = regr.predict(xval)
print(yval)

xval = np.linspace(0, 2, 20).reshape(-1, 1)
yval = regr.predict(xval)
plt.plot(xval, yval, color='black')
plt.scatter(xp, yp)
plt.show()
