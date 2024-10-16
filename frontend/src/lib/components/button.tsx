import { ComponentProps } from "react";
import { twMerge } from "tailwind-merge";
import { cva, VariantProps } from "class-variance-authority";

const buttonStyles = cva(
  "py-1 px-6 active:scale-95 disabled:opacity-50 border border-transparent disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:pointer-events-none",
  {
    variants: {
      variant: {
        default: "bg-indigo-600 text-white hover:bg-indigo-700",
        outline: "bg-white text-indigo-600 hover:bg-gray-50 border-indigo-600",
        ghost: "bg-white text-indigo-600 hover:bg-gray-50",
      },
      size: {
        sm: "text-sm",
        md: "text-base",
        lg: "text-lg",
      },
      hasIcon: {
        true: "flex items-center justify-center space-x-2 px-1 py-1",
        false: "",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "md",
      hasIcon: false,
    },
  }
);

type ButtonProps = ComponentProps<"button"> & VariantProps<typeof buttonStyles>;

export const Button = ({ className, variant, size, ...props }: ButtonProps) => {
  return (
    <button
      className={twMerge(buttonStyles({ variant, size }), className)}
      {...props}
    />
  );
};
