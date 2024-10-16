import { EventsOn } from "../../../wailsjs/runtime/runtime";
import { useEffect, useMemo, useState } from "react";
import { GameStateSpectating, AllPlayers } from "csgo-gsi-types";
import { GetSteamAvatar } from "../../../wailsjs/go/main/App";

export function useGameState() {
  const [gameState, setGameState] = useState<GameStateSpectating | null>(null);

  useEffect(() => {
    EventsOn("gsi-raw-state", (data: any) => {
      const parsed = typeof data === "string" ? JSON.parse(data) : data;
      setGameState(parsed);
    });
  }, []);

  return { gameState };
}

export type Player = __GSICSGO.PlayerList & {
  _id: string;
};

export type UseGameStateReturnType = {
  teamCT: Player[];
  teamT: Player[];
  avatars: Record<string, string>;
};

export function useGameStateTeams(): UseGameStateReturnType {
  const { gameState } = useGameState();

  const [trackedIds, setTrackedIds] = useState<string[] | null>(null);

  const [avatars, setAvatars] = useState<Record<string, string>>({});

  const teamT = useMemo(() => {
    if (gameState?.allplayers) {
      const allPlayers = Object.keys(gameState.allplayers).map((key) => ({
        ...gameState.allplayers[key],
        _id: key,
      }));

      return allPlayers.filter((p) => p.team === "T");
    }
    return [];
  }, [gameState?.allplayers]);

  const teamCT = useMemo(() => {
    if (gameState?.allplayers) {
      const allPlayers = Object.keys(gameState.allplayers).map((key) => ({
        ...gameState.allplayers[key],
        _id: key,
      }));

      return allPlayers.filter((p) => p.team === "CT");
    }
    return [];
  }, [gameState?.allplayers]);

  useEffect(() => {
    const updated = [...teamCT, ...teamT]
      .map((i) => i._id)
      .sort((a, b) => a.localeCompare(b));

    if (!trackedIds) {
      setTrackedIds(updated);
      return;
    }

    const allMatch =
      trackedIds.every((id, index) => id === updated[index]) &&
      trackedIds.length !== 0;

    if (allMatch) {
      return;
    }

    setTrackedIds(updated);
  }, [teamCT, teamT]);

  useEffect(() => {
    if (!trackedIds) {
      return;
    }

    const fetchAvatars = async () => {
      const avatars: { [key: string]: string } = {};

      for (const id of trackedIds) {
        const avatarFullUrl = await GetSteamAvatar(id);
        console.log(avatarFullUrl);
        avatars[id] = avatarFullUrl;
      }

      setAvatars(avatars);
    };

    fetchAvatars();
  }, [trackedIds]);

  return {
    teamCT,
    teamT,
    avatars,
  };
}
