import numpy as np
import re

def parse_equation(equation):
    """Parses a linear equation of the form ax + by + cz = d and returns coefficients a, b, c and d."""
    # Remove spaces for easier processing
    equation = equation.replace(' ', '')
    
    # Split the equation into left-hand side (LHS) and right-hand side (RHS)
    if '=' not in equation:
        raise ValueError("Equation must contain '=' sign")
    
    lhs, rhs = equation.split('=')
    
    # Extract the coefficients for x, y, z
    coefficients = [0, 0, 0]
    terms = re.findall(r'[+-]?\d*x|[+-]?\d*y|[+-]?\d*z', lhs)
    
    for term in terms:
        if 'x' in term:
            if term in ['x', '+x']:
                coefficients[0] += 1
            elif term == '-x':
                coefficients[0] -= 1
            else:
                coefficients[0] += int(term.replace('x', ''))
        elif 'y' in term:
            if term in ['y', '+y']:
                coefficients[1] += 1
            elif term == '-y':
                coefficients[1] -= 1
            else:
                coefficients[1] += int(term.replace('y', ''))
        elif 'z' in term:
            if term in ['z', '+z']:
                coefficients[2] += 1
            elif term == '-z':
                coefficients[2] -= 1
            else:
                coefficients[2] += int(term.replace('z', ''))
    
    # Handle constant term on RHS
    constant = int(rhs)
    coefficients.append(constant)
    
    return coefficients

def main():
    equations = []
    for i in range(3):
        eq = input(f"Enter equation {i+1} (in the form ax + by + cz = d): ")
        equations.append(eq)
    
    A = []
    B = []

    for eq in equations:
        coefficients = parse_equation(eq)
        A.append(coefficients[:3])
        B.append(coefficients[3])

    A = np.array(A, dtype=float)
    B = np.array(B, dtype=float)
    
    print("\nMatrix A:")
    print(A)
    
    print("\nMatrix B:")
    print(B)
    
    try:
        X = np.linalg.solve(A, B)
        print("\nSolution for X:")
        print(f"x = {X[0]:.2f}")
        print(f"y = {X[1]:.2f}")
        print(f"z = {X[2]:.2f}")
    except np.linalg.LinAlgError:
        print("The system of equations does not have a unique solution.")

if __name__ == "__main__":
    main()

