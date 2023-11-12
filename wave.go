package main

type EnemySpawnInfo struct {
	SpawnFrame int // 何フレーム後に敵をスポーンさせるか
}

type Wave struct {
	EnemySpawns []EnemySpawnInfo // このウェーブでの敵の出現情報
	TotalFrames int              // このウェーブの持続フレーム数（次のウェーブが開始するまでのフレーム数）
}

type Stage struct {
	Waves []Wave // このステージにおける各ウェーブの情報
}

// サンプルの敵出現情報
var sampleStage = Stage{
	Waves: []Wave{
		{
			EnemySpawns: []EnemySpawnInfo{
				{SpawnFrame: 60},  // 1秒後 (60fps 前提)
				{SpawnFrame: 120}, // 2秒後
				{SpawnFrame: 180}, // 3秒後
			},
			TotalFrames: 300, // 5秒間のウェーブ
		},
		{
			EnemySpawns: []EnemySpawnInfo{
				{SpawnFrame: 60},
				{SpawnFrame: 90},
				{SpawnFrame: 150},
				{SpawnFrame: 210},
			},
			TotalFrames: 360, // 6秒間のウェーブ
		},
		{
			EnemySpawns: []EnemySpawnInfo{
				{SpawnFrame: 60},
				{SpawnFrame: 90},
				{SpawnFrame: 120},
				{SpawnFrame: 150},
				{SpawnFrame: 180},
				{SpawnFrame: 210},
			},
			TotalFrames: 360, // 6秒間のウェーブ
		},
	},
}

var debugStage = Stage{
	Waves: []Wave{
		{
			EnemySpawns: []EnemySpawnInfo{
				{SpawnFrame: 60}, // 1秒後 (60fps 前提)
			},
			TotalFrames: 300, // 5秒間のウェーブ
		},
	},
}
