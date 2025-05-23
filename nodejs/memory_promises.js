const os = require('os');

function simulateTask(id) {
  return new Promise((resolve) => {
    setTimeout(resolve, 100); // Simula tarefa
  });
}

async function main() {
  const numTasks = 1_000_000;
  const before = process.memoryUsage().heapUsed / 1024 / 1024;

  const tasks = [];
  for (let i = 0; i < numTasks; i++) {
    tasks.push(simulateTask(i));
  }

  await Promise.all(tasks);
  const after = process.memoryUsage().heapUsed / 1024 / 1024;

  console.log(`Total tasks: ${numTasks}`);
  console.log(`Memory used: ${(after - before).toFixed(2)} MB`);
}

main();
