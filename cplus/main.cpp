#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <regex>
#include <Eigen/Dense>

using namespace std;
using namespace Eigen;

vector<double> parse_equation(const string &equation) {
    vector<double> coefficients(4, 0.0);
    regex term_regex("([+-]?\\d*)\\s*([xyz])");
    sregex_iterator iter(equation.begin(), equation.end(), term_regex);
    sregex_iterator end;
    
    for (; iter != end; ++iter) {
        string coefficient_str = (*iter)[1].str();
        char variable = (*iter)[2].str()[0];
        double coefficient = (coefficient_str.empty() || coefficient_str == "+" || coefficient_str == "-") ? 
                             (coefficient_str == "-" ? -1 : 1) : 
                             stod(coefficient_str);
        switch (variable) {
            case 'x': coefficients[0] += coefficient; break;
            case 'y': coefficients[1] += coefficient; break;
            case 'z': coefficients[2] += coefficient; break;
        }
    }
    
    size_t equals_pos = equation.find('=');
    if (equals_pos != string::npos) {
        coefficients[3] = stod(equation.substr(equals_pos + 1));
    } else {
        throw invalid_argument("Equation must contain '=' sign");
    }

    return coefficients;
}

int main() {
    vector<string> equations(3);
    cout << "Enter 3 equations in the form 'ax + by + cz = d':\n";
    for (int i = 0; i < 3; ++i) {
        cout << "Enter equation " << i + 1 << ": ";
        getline(cin, equations[i]);
    }
    
    Matrix3d A;
    Vector3d B;
    for (int i = 0; i < 3; ++i) {
        vector<double> coefficients = parse_equation(equations[i]);
        A(i, 0) = coefficients[0];
        A(i, 1) = coefficients[1];
        A(i, 2) = coefficients[2];
        B(i) = coefficients[3];
    }

    cout << "\nMatrix A:\n" << A << "\n";
    cout << "\nVector B:\n" << B << "\n";

    Vector3d X;
    try {
        X = A.colPivHouseholderQr().solve(B);
        cout << "\nSolution for X:\n";
        cout << "x = " << X(0) << "\n";
        cout << "y = " << X(1) << "\n";
        cout << "z = " << X(2) << "\n";
    } catch (const exception &e) {
        cout << "The system of equations does not have a unique solution.\n";
    }

    return 0;
}