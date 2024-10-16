import { OpenSelf } from "../../wailsjs/go/main/App";
import { Button } from "../lib/components/button";
import { useGameState } from "../lib/hooks/useGameState";

export default function Overview() {
  const { gameState } = useGameState();
  return (
    <div className="p-4 flex justify-center items-center flex-col min-h-screen">
      <h2 className="font-medium">Waiting for game..</h2>

      <Button
        onClick={() => {
          OpenSelf("hud");
        }}
        className="mt-4"
        size="lg"
      >
        Start Game
      </Button>

      <pre>
        <code>{JSON.stringify(gameState, null, 2)}</code>
      </pre>
    </div>
  );
}
