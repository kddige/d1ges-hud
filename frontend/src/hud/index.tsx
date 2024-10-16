import { Crosshair, Skull } from "lucide-react";
import { FC, useEffect, useMemo } from "react";
import {
  Player,
  UseGameStateReturnType,
  useGameState,
  useGameStateTeams,
} from "../lib/hooks/useGameState";
import { onGameEvent } from "../lib/events/gameEvents";
import { EventsOff, EventsOn } from "../../wailsjs/runtime/runtime";



export default function Hud() {
  const { gameState } = useGameState();

  const teams = useGameStateTeams();

  useEffect(() => {
    const cancel = onGameEvent("GameMap", console.log);

    return () => {
      cancel();
    }
  }, []);

  useEffect(() => {
  }, [gameState]);

  return (
    <div className="h-screen py-4 px-8 flex flex-col">
      <div className="flex flex-row gap-8">
        <div className="basis-96">
          <PhaseBar
            map={gameState?.map?.name ?? ""}
            phase={gameState?.map?.phase ?? ""}
          />
        </div>
        <ScoreBar
          t={{
            score: gameState?.map?.team_t?.score ?? 0,
          }}
          ct={{
            score: gameState?.map?.team_ct?.score ?? 0,
          }}
          phase_ends_in={gameState?.phase_countdowns?.phase_ends_in ?? 0}
        />
        <div className="basis-96"></div>
      </div>
      <div className="flex flex-1 items-center justify-between">
        <PlayerList avatars={teams.avatars} players={teams.teamCT} team="ct" />
        <PlayerList
          avatars={teams.avatars}
          players={teams.teamT}
          team="t"
          reversed
        />
      </div>
    </div>
  );
}

type PhaseBarProps = {
  map: string;
  phase: string;
};
const PhaseBar: FC<PhaseBarProps> = ({ map, phase }) => {
  return (
    <div className="bg-black rounded-xl px-4 py-1 flex">
      <p className="font-bold">
        <span className="text-white/50">Playing</span> {map}
      </p>
      <div className="w-[1px] bg-white/50 mx-2 my-1"></div>
      <p className="font-bold">
        <span className="text-white/50">Phase</span> {phase}
      </p>
    </div>
  );
};

type ScoreBarProps = {
  phase_ends_in: number;
  ct: {
    score: number;
  };
  t: {
    score: number;
  };
};
const ScoreBar: FC<ScoreBarProps> = ({ phase_ends_in, ct, t }) => {
  const phaseEndInPretty = useMemo(() => {
    if (phase_ends_in < 0) return "00:00";
    return new Date((phase_ends_in ?? 0) * 1000).toISOString().substr(14, 5);
  }, [phase_ends_in]);

  return (
    <div className="flex-1 bg-black/80 rounded-lg flex py-4 px-4">
      <ScoreBarTeam team="ct" score={ct.score} />
      <div className="basis-24 text-center">
        <p className="text-3xl font-bold">{phaseEndInPretty}</p>
      </div>
      <ScoreBarTeam team="t" score={t.score} reverse />
    </div>
  );
};

type ScoreBarTeamProps = {
  team: "ct" | "t";
  score: number;
  reverse?: boolean;
};
const ScoreBarTeam: FC<ScoreBarTeamProps> = ({ team, score, reverse }) => {
  const teamBg = team === "ct" ? "bg-blue-400" : "bg-yellow-400";
  let teamScoreContainerClass =
    team === "ct"
      ? "border-blue-400 text-blue-400 pr-3 border-r-4"
      : "border-yellow-400 text-yellow-400";

  if (reverse) {
    teamScoreContainerClass += " border-l-4 border-r-0 pl-3 pr-0";
  }

  return (
    <div
      className={`flex-1 flex justify-between ${
        reverse ? "flex-row-reverse" : "flex-row"
      }`}
    >
      <div className={`${teamBg} h-full aspect-square rounded-full`}></div>
      <p className="text-3xl font-bold">
        {team === "ct" ? "Counter-Terrorists" : "Terrorists"}
      </p>
      <div className={`${teamScoreContainerClass}`}>
        <p className="text-3xl font-bold">{score}</p>
      </div>
    </div>
  );
};

type PlayerListProps = {
  reversed?: boolean;
  team: "ct" | "t";
  players: UseGameStateReturnType["teamCT"] | UseGameStateReturnType["teamT"];
  avatars: Record<string, string>;
};
const PlayerList: FC<PlayerListProps> = ({
  reversed,
  team,
  players,
  avatars,
}) => {
  const wrapperClass = `basis-96 space-y-2 ${reversed ? "-mr-8" : "-ml-8"}`;
  return (
    <div className={wrapperClass}>
      {players.map((p) => (
        <PlayerListItem
          avatar={avatars[p._id]}
          key={p._id}
          player={p}
          team={team}
          reversed={reversed}
        />
      ))}
    </div>
  );
};

type PlayerListItemProps = {
  reversed?: boolean;
  team: "ct" | "t";
  player: Player;
  avatar?: string;
};
const PlayerListItem: FC<PlayerListItemProps> = ({
  reversed,
  team,
  player,
  avatar,
}) => {
  const wrapperClass = `bg-black/80 w-full flex p-2 gap-4 ${
    reversed ? "flex-row-reverse justify-start" : "flex-row"
  }`;

  const nameContainer = `flex w-full justify-between ${
    reversed ? "flex-row-reverse" : ""
  }`;

  const iconColorClass = team === "ct" ? "text-blue-400" : "text-yellow-400";

  const contentClass = `flex gap-2 ${reversed ? "justify-end" : ""}`;
  return (
    <div className={wrapperClass}>
      <div
        style={
          avatar
            ? { backgroundImage: `url(${avatar})`, backgroundSize: "contain" }
            : {}
        }
        className="rounded-full bg-gray-400 h-14 w-14 aspect-square"
      ></div>
      <div className="space-y-1 w-full">
        <div className={nameContainer}>
          <p>{player.name}</p>
          <div className={`flex gap-4 ${reversed ? "flex-row-reverse" : ""}`}>
            <div
              className={`flex items-center gap-1 w-10 justify-end ${
                reversed ? "flex-row-reverse" : ""
              }`}
            >
              <span>{player.match_stats.deaths}</span>
              <Skull className={`h-4 w-4 ${iconColorClass}`} />
            </div>
            <div
              className={`flex items-center gap-1 w-10 justify-end ${
                reversed ? "flex-row-reverse" : ""
              }`}
            >
              <span>{player.match_stats.kills}</span>
              <Crosshair className={`h-4 w-4 ${iconColorClass}`} />
            </div>
          </div>
        </div>
        <div className={contentClass}></div>
      </div>
    </div>
  );
};
