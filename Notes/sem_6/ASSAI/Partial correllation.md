# Partial Correlation

**Partial correlation**, calculates the [correlation](https://datatab.net/tutorial/correlation) between two variables, while excluding the effect of a third variable. This makes it possible to find out whether the correlation _rxy_ between variables _x_ and _y_ is produced by the variable _z_.

![Partial Correlation](https://datatab.net/assets/tutorial/partialkorrelation.png "Partial Correlation")

The partial correlation _rxy,z_ tells how strongly the variable _x_ correlates with the variable _y_, if the [correlation](https://datatab.net/tutorial/correlation) of both variables with the variable _z_ is calculated out.

## Calculate Partial Correlation

For the calculation of the partial correlation, the three [correlations](https://datatab.net/tutorial/correlation) between the individual variables are required. The partial correlation then results in

![Partial Correlation Equation](https://datatab.net/assets/tutorial/equ/partialkorrelation_gleichung.svg "Partial Correlation Equation")

- _rxy_ = Correlation between variable _x_ and _y_
- _rxz_ = Correlation of the third variable _z_ with the variable _x_
- _ryz_ = Correlation of the third variable _z_ with the variable _y_