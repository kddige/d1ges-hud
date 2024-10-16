import { PropsWithChildren, FC, ComponentProps } from "react";
import { twMerge } from "tailwind-merge";
import { cva, VariantProps } from "class-variance-authority";
import { NavLink } from "react-router-dom";

type SidebarProps = {
  title: string;
  subtitle: string;
} & PropsWithChildren;
const Sidebar: FC<SidebarProps> = ({ title, subtitle, children }) => {
  return (
    <div className="basis-64 from-zinc-900 to-indigo-900 bg-gradient-to-b py-8 shadow-lg">
      <div className="px-4">
        <h1>{title}</h1>
        <small className="ml-12">{subtitle}</small>
      </div>

      <div className="mt-4 w-full">{children}</div>
    </div>
  );
};

const sidebarItemStyles = cva(
  "p-4 text-start w-full flex gap-2 uppercase font-medium items-center",
  {
    variants: {
      active: {
        true: "bg-indigo-600 text-white",
        false: "hover:bg-white/10",
      },
    },
    defaultVariants: {
      active: false,
    },
  }
);

type SidebarItemProps = ComponentProps<typeof NavLink> &
  VariantProps<typeof sidebarItemStyles> & {
    className?: string;
  };
const SidebarItem: FC<SidebarItemProps> = ({
  children,
  className,
  active,
  ...props
}) => {
  return (
    <NavLink
      className={({ isActive }) =>
        twMerge(sidebarItemStyles({ active: active ?? isActive }), className)
      }
      {...props}
    >
      {children}
    </NavLink>
  );
};

export { Sidebar, SidebarItem };
