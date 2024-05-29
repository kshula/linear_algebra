package linear;
import java.util.Scanner;
import org.apache.commons.math3.linear.*;

public class SolveLinearEquations {
    
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        double[][] coefficients = new double[3][3];
        double[] constants = new double[3];

        System.out.println("Enter 3 equations in the form 'ax + by + cz = d':");

        for (int i = 0; i < 3; i++) {
            System.out.print("Enter equation " + (i + 1) + ": ");
            String equation = scanner.nextLine();
            double[] parsedEquation = parseEquation(equation);
            System.arraycopy(parsedEquation, 0, coefficients[i], 0, 3);
            constants[i] = parsedEquation[3];
        }

        RealMatrix A = MatrixUtils.createRealMatrix(coefficients);
        RealVector B = MatrixUtils.createRealVector(constants);

        System.out.println("\nMatrix A:");
        System.out.println(A);

        System.out.println("\nVector B:");
        System.out.println(B);

        try {
            DecompositionSolver solver = new LUDecomposition(A).getSolver();
            RealVector solution = solver.solve(B);

            System.out.println("\nSolution for X:");
            System.out.printf("x = %.2f\n", solution.getEntry(0));
            System.out.printf("y = %.2f\n", solution.getEntry(1));
            System.out.printf("z = %.2f\n", solution.getEntry(2));
        } catch (SingularMatrixException e) {
            System.out.println("The system of equations does not have a unique solution.");
        }

        scanner.close();
    }

    private static double[] parseEquation(String equation) {
        double[] coefficients = new double[4]; // a, b, c, and d
        String lhs = equation.split("=")[0].replaceAll("\\s+", "");
        String rhs = equation.split("=")[1].replaceAll("\\s+", "");
        coefficients[3] = Double.parseDouble(rhs);

        String[] terms = lhs.split("(?=[+-])");

        for (String term : terms) {
            if (term.contains("x")) {
                coefficients[0] = term.equals("x") || term.equals("+x") ? 1 :
                                  term.equals("-x") ? -1 : Double.parseDouble(term.replace("x", ""));
            } else if (term.contains("y")) {
                coefficients[1] = term.equals("y") || term.equals("+y") ? 1 :
                                  term.equals("-y") ? -1 : Double.parseDouble(term.replace("y", ""));
            } else if (term.contains("z")) {
                coefficients[2] = term.equals("z") || term.equals("+z") ? 1 :
                                  term.equals("-z") ? -1 : Double.parseDouble(term.replace("z", ""));
            }
        }

        return coefficients;
    }
}