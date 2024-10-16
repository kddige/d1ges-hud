import { Outlet } from "react-router-dom";
import { HomeIcon, SettingsIcon } from "lucide-react";
import { Sidebar, SidebarItem } from "../lib/components/sidebar";
import { GetMode } from "../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import Hud from "../hud";

export default function Root() {
  const [mode, setMode] = useState<string | null>(null);

  useEffect(() => {
    GetMode().then((mode) => {
      setMode(mode);
      document.body.classList.add("mode-" + mode);
    });
  }, []);

  if (mode === null) {
    return null;
  }

  if (mode === "hud") {
    return <Hud />;
  }

  // default / normal
  return (
    <div className="h-screen flex flex-row">
      <pre>
        <code>{JSON.stringify(mode, null, 2)}</code>
      </pre>
      <Sidebar title="DHM" subtitle="D1ges Hud manager">
        <SidebarItem to={"/"}>
          <HomeIcon className="h-5 w-5" />
          Overview
        </SidebarItem>
        <SidebarItem to={"/settings"}>
          <SettingsIcon className="h-5 w-5" />
          Settings
        </SidebarItem>
      </Sidebar>
      <main className="flex-1 overflow-y-auto">
        <Outlet />
        <span className="absolute right-0 bottom-0 backdrop-invert px-1 text-inherit">
          v0.0.1-alpha
        </span>
      </main>
    </div>
  );
}
