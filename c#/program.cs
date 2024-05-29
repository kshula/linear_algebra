using System;
using MathNet.Numerics.LinearAlgebra;

namespace SolveLinearEquations
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Enter 3 equations in the form 'ax + by + cz = d':");

            var A = Matrix<double>.Build.Dense(3, 3, 0.0);
            var B = Vector<double>.Build.Dense(3, 0.0);

            for (int i = 0; i < 3; i++)
            {
                Console.Write($"Enter equation {i + 1}: ");
                string equation = Console.ReadLine();
                var parsedEquation = ParseEquation(equation);
                A.SetRow(i, new double[] { parsedEquation[0], parsedEquation[1], parsedEquation[2] });
                B[i] = parsedEquation[3];
            }

            Console.WriteLine("\nMatrix A:");
            Console.WriteLine(A);

            Console.WriteLine("\nVector B:");
            Console.WriteLine(B);

            try
            {
                var X = A.Solve(B);
                Console.WriteLine("\nSolution for X:");
                Console.WriteLine($"x = {X[0]:F2}");
                Console.WriteLine($"y = {X[1]:F2}");
                Console.WriteLine($"z = {X[2]:F2}");
            }
            catch (ArgumentException)
            {
                Console.WriteLine("The system of equations does not have a unique solution.");
            }
        }

        static double[] ParseEquation(string equation)
        {
            double[] coeffs = new double[4]; // a, b, c, and d

            string[] parts = equation.Split('=');
            if (parts.Length != 2)
                throw new ArgumentException("Equation must contain '=' sign");

            string lhs = parts[0].Trim();
            string rhs = parts[1].Trim();

            string[] terms = lhs.Split(new char[] { '+', '-' }, StringSplitOptions.RemoveEmptyEntries);

            foreach (var term in terms)
            {
                string t = term.Trim();
                if (t.EndsWith("x"))
                    coeffs[0] += ParseCoefficient(t.Substring(0, t.Length - 1));
                else if (t.EndsWith("y"))
                    coeffs[1] += ParseCoefficient(t.Substring(0, t.Length - 1));
                else if (t.EndsWith("z"))
                    coeffs[2] += ParseCoefficient(t.Substring(0, t.Length - 1));
                else
                    coeffs[3] = double.Parse(t); // constant term
            }

            return coeffs;
        }

        static double ParseCoefficient(string coefficient)
        {
            if (string.IsNullOrWhiteSpace(coefficient))
                return 1.0;

            if (coefficient == "-")
                return -1.0;

            return double.Parse(coefficient);
        }
    }
}