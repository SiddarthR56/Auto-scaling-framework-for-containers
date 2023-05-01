import numpy as np
import tensorflow as tf
import gym

# Define the environment for the container scaling problem
class ContainerEnv(gym.Env):
    def __init__(self, max_containers, max_cpu, max_memory, max_network):
        self.max_containers = max_containers
        self.max_cpu = max_cpu
        self.max_memory = max_memory
        self.max_network = max_network
        self.action_space = gym.spaces.Discrete(max_containers + 1)
        self.observation_space = gym.spaces.Box(low=0, high=1, shape=(4,), dtype=np.float32)
        self.state = np.zeros((4,), dtype=np.float32)
        
    def reset(self):
        self.state = np.zeros((4,), dtype=np.float32)
        return self.state
        
    def step(self, action):
        # Apply the action to the system and update the state
        if action > self.state[0]:
            # Add containers
            self.state[1] += (action - self.state[0]) * 0.1
            self.state[2] += (action - self.state[0]) * 0.2
            self.state[3] += (action - self.state[0]) * 0.3
        elif action < self.state[0]:
            # Remove containers
            self.state[1] -= (self.state[0] - action) * 0.1
            self.state[2] -= (self.state[0] - action) * 0.2
            self.state[3] -= (self.state[0] - action) * 0.3
            
        self.state[0] = action
        
        # Compute the reward based on the current state
        reward = -0.1 * self.state[1] - 0.2 * self.state[2] - 0.3 * self.state[3]
        
        # Check if the state is valid
        done = False
        if self.state[1] > self.max_cpu or self.state[2] > self.max_memory or self.state[3] > self.max_network:
            done = True
            reward = -1
            
        # Return the new state, reward, and done flag
        return self.state, reward, done, {}
        
        
# Define the DQN agent
class DQNAgent:
    def __init__(self, env):
        self.env = env
        self.gamma = 0.99
        self.epsilon = 1.0
        self.epsilon_min = 0.01
        self.epsilon_decay = 0.999
        self.learning_rate = 0.001
        self.memory = []
        self.model = self.build_model()
        
    def build_model(self):
        model = tf.keras.models.Sequential([
            tf.keras.layers.Dense(32, activation='relu', input_shape=(4,)),
            tf.keras.layers.Dense(32, activation='relu'),
            tf.keras.layers.Dense(self.env.action_space.n, activation='linear')
        ])
        model.compile(loss='mse', optimizer=tf.keras.optimizers.Adam(lr=self.learning_rate))
        return model
        
    def act(self, state):
        if np.random.rand() < self.epsilon:
            return self.env.action_space.sample()
        q_values = self.model.predict(state)
        return np.argmax(q_values[0])
        
    def remember(self, state, action, reward, next_state, done):
        self.memory.append((state, action, reward, next_state, done))
        
    def replay(self, batch_size):
        if len(self.memory)
